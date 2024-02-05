// Package collectors contains custom prometheus collectors that are not just simple instrumented metrics
package economic

import (
	"errors"
	"strconv"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/bls"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

// Collector exports the economic metrics listed at
//
//	https://eigen.nethermind.io/docs/spec/metrics/metrics-examples#economics-metrics
//
// these are metrics that are exported not via instrumentation, but instead by proxying
// a call to the relevant eigenlayer contracts
// Collector should be registered in the same prometheus registry that is passed to metrics.NewEigenMetrics
// so that they are exported on the same port
type Collector struct {
	// TODO(samlaf): we use a chain as the backend for now, but should eventually move to a subgraph
	elReader          elcontracts.ELReader
	avsRegistryReader avsblsregistry.AvsRegistryReader
	logger            logging.Logger
	// params to query the metrics for
	operatorAddr common.Address
	operatorId   types.BlsOperatorId
	quorumNames  map[types.QuorumNum]string
	// metrics
	// TODO(samlaf): I feel like eigenlayer-core metrics like slashingStatus and delegatedShares, which are not avs specific,
	// should not be here, and should instead be collected by some eigenlayer-cli daemon or something, since
	// otherwise every avs will be exporting these same metrics for no reason.

	// slashingStatus is 1 when the operator has been slashed, and 0 otherwise
	slashingStatus *prometheus.Desc
	// registeredStake is the current stake of the operator in the avs registry contract's view
	// stake is a weighted linear combination of strategyManager shares in different strategies
	// stakes are updated by calling updateStakes() on the StakeRegistry contract, which pulls the
	// current strategy shares from the strategyManager contracts
	// eg. suppose the avs has defined an ethLST quorum consisting of stETH, rETH, and cbETH, with each weighted 1
	// (a decentralization focused avs might want to give more weight to rETH for eg, but we'll keep it simple)
	// if the operator has 100 stETH, 200 rETH, and 300 cbETH delegated in the strategyManager, then it's
	// registeredStake on the StakeRegistry contract will be 600. If a staker queues a withdrawal of its 300 cbETH,
	// then the operator's registeredStake will remain 600 until updateStakes() is called, at which point it will
	// drop to the correct value of 300.
	registeredStake *prometheus.Desc

	// TODO(samlaf): Removing this as not part of avs node spec anymore.
	// delegatedShares is the total shares delegated to the operator in each strategy
	// strategies are currently token specific, and we have one for cbETH, rETH, and stETH,
	// but they will eventually be more generic and could support multiple tokens
	// right now, 1 token deposit = 1 share
	// delegatedShares *prometheus.Desc
}

var _ prometheus.Collector = (*Collector)(nil)

func NewCollector(
	elReader elcontracts.ELReader, avsRegistryReader avsblsregistry.AvsRegistryReader,
	avsName string, logger logging.Logger,
	operatorAddr common.Address, quorumNames map[types.QuorumNum]string,
) *Collector {
	return &Collector{
		elReader:          elReader,
		avsRegistryReader: avsRegistryReader,
		logger:            logger,
		operatorAddr:      operatorAddr,
		// we don't fetch operatorId here because operator might not yet be registered (and hence not have an operatorId)
		// we cache operatorId dynamically in the collect function() inside, which allows constructing collector before registering operator
		operatorId:  [32]byte{},
		quorumNames: quorumNames,
		slashingStatus: prometheus.NewDesc(
			types.EigenPromNamespace+"_slashing_status",
			"Whether the operator has been slashed",
			// TODO(samlaf): would be nice to have avs_name as a variable label
			[]string{},
			// no avs_name label since slashing is not avs specific
			prometheus.Labels{},
		),
		registeredStake: prometheus.NewDesc(
			types.EigenPromNamespace+"_registered_stakes",
			"Operator stake in <quorum> of <avs_name>'s StakeRegistry contract",
			[]string{"quorum_number", "quorum_name"},
			prometheus.Labels{"avs_name": avsName},
		),
		// delegatedShares: prometheus.NewDesc(
		// 	"eigen_delegated_shares",
		// 	"Operator delegated shares in <strategy>, in units of <unit> (wei, gwei, or eth).",
		// 	// each strategy has a single token for now (eg. stETH, rETH, cbETH), so we add the token label for now
		// 	[]string{"strategy", "unit", "token"},
		// 	prometheus.Labels{},
		// ),
	}
}

// Describe function for the exported slashingIncurredTotal and balanceTotal metrics
// see https://github.com/prometheus/client_golang/blob/v1.16.0/prometheus/collector.go#L27
// Describe is implemented with DescribeByCollect. That's possible because the
// Collect method will always return the same two metrics with the same two
// descriptors.
func (ec *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- ec.slashingStatus
	ch <- ec.registeredStake
	// ch <- ec.delegatedShares
}

// initialize the operatorId if it hasn't already been initialized
func (ec *Collector) initOperatorId() error {
	if ec.operatorId == [32]byte{} {
		operatorId, err := ec.avsRegistryReader.GetOperatorId(&bind.CallOpts{}, ec.operatorAddr)
		if err != nil {
			ec.logger.Error("Failed to get operator id", "err", err)
			return err
		}
		if operatorId == [32]byte{} {
			return errors.New("operator not registered")
		}
		ec.operatorId = operatorId
	}
	return nil
}

// Collect function for the exported slashingIncurredTotal and balanceTotal metrics
// see https://github.com/prometheus/client_golang/blob/v1.16.0/prometheus/collector.go#L27
// collect just makes jsonrpc calls to the slasher and delegationManager and then creates
// constant metrics with the results
func (ec *Collector) Collect(ch chan<- prometheus.Metric) {
	// collect slashingStatus metric
	// TODO(samlaf): note that this call is not avs specific, so every avs will have the same value if the operator has been slashed
	// if we want instead to only output 1 if the operator has been slashed for a specific avs, we have 2 choices:
	// 1. keep this collector format but query the OperatorFrozen event from a subgraph
	// 2. subscribe to the event and keep a local state of whether the operator has been slashed, exporting it via normal prometheus instrumentation
	operatorIsFrozen, err := ec.elReader.OperatorIsFrozen(nil, ec.operatorAddr)
	if err != nil {
		ec.logger.Error("Failed to get slashing incurred", "err", err)
	} else {
		operatorIsFrozenFloat := 0.0
		if operatorIsFrozen {
			operatorIsFrozenFloat = 1.0
		}
		ch <- prometheus.MustNewConstMetric(ec.slashingStatus, prometheus.CounterValue, operatorIsFrozenFloat)
	}

	// collect registeredStake metric
	err = ec.initOperatorId()
	if err != nil {
		ec.logger.Warn("Failed to fetch and cache operator id. Skipping collection of registeredStake metric.", "err", err)
	} else {
		// probably should start using the avsregistry service instead of avsRegistryReader so that we can
		// swap out backend for a subgraph eventually
		quorumStakeMap, err := ec.avsRegistryReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, ec.operatorId)
		if err != nil {
			ec.logger.Error("Failed to get operator stake", "err", err)
		} else {
			for quorumNum, stake := range quorumStakeMap {
				stakeFloat64, _ := stake.Float64()
				ch <- prometheus.MustNewConstMetric(
					ec.registeredStake, prometheus.GaugeValue, stakeFloat64, strconv.Itoa(int(quorumNum)), ec.quorumNames[quorumNum],
				)
			}
		}
	}

	// collect delegatedShares metric
	// for _, strategyAddr := range ec.strategyAddrs {
	// 	sharesWei, err := ec.elReader.GetOperatorSharesInStrategy(context.Background(), ec.operatorAddr, strategyAddr)
	// 	if err != nil {
	// 		ec.logger.Error("Failed to get shares", "err", err)
	// 	} else {
	// 		// We'll emit all 3 units in case this is needed for whatever reason
	// 		// might want to change this behavior if this is emitting too many metrics
	// 		sharesWeiFloat, _ := sharesWei.Float64()
	// 		// TODO(samlaf): add the token name.. probably need to have a hardcoded dict per env (mainnet, goerli, etc)? Is it really that important..?
	// 		ch <- prometheus.MustNewConstMetric(ec.delegatedShares, prometheus.GaugeValue, sharesWeiFloat, strategyAddr.String(), "wei", "token")

	// 		sharesGweiFloat, _ := sharesWei.Div(sharesWei, big.NewInt(1e9)).Float64()
	// 		ch <- prometheus.MustNewConstMetric(ec.delegatedShares, prometheus.GaugeValue, sharesGweiFloat, strategyAddr.String(), "gwei", "token")

	// 		sharesEtherFloat, _ := sharesWei.Div(sharesWei, big.NewInt(1e18)).Float64()
	// 		ch <- prometheus.MustNewConstMetric(ec.delegatedShares, prometheus.GaugeValue, sharesEtherFloat, strategyAddr.String(), "ether", "token")
	// 	}
	// }
}
