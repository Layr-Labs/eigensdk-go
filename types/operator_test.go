package types

import (
	"errors"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/stretchr/testify/assert"
)

func TestOperatorValidate(t *testing.T) {
	var tests = []struct {
		name                    string
		operator                Operator
		metadataURLresponseCode int
		operatorMetadata        OperatorMetadata
		wantErr                 bool
		expectedErr             error
	}{
		{
			name: "successful operator validation",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			},
			wantErr: false,
		},
		{
			name: "valid operator validation - ZeroAddress delegation approver address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: ZeroAddress,
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			},
			wantErr: false,
		},
		{
			name: "failed operator validation - empty metadata url",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrEmptyUrl),
		},
		{
			name: "failed operator validation - localhost metadata url",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "http://localhost:8080/metadata.json",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrUrlPointingToLocalServer),
		},
		{
			name: "failed operator validation - 127.0.0.1 metadata url",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "http://127.0.0.1:8080/metadata.json",
			},
			wantErr:     true,
			expectedErr: utils.WrapError(ErrInvalidMetadataUrl, utils.ErrUrlPointingToLocalServer),
		},
		{
			name: "failed operator validation - bad metadata",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://example.com/metadata.json",
			},
			wantErr: true,
			expectedErr: utils.WrapError(
				ErrReadingMetadataUrlResponse,
				errors.New("error fetching url: 404 Not Found"),
			),
		},
		{
			name: "failed operator validation - wrong operator address",
			operator: Operator{
				Address:                   "0xa",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://example.com/metadata.json",
			},
			wantErr:     true,
			expectedErr: ErrInvalidOperatorAddress,
		},
		{
			name: "failed operator validation - wrong earning receivers address address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				EarningsReceiverAddress:   "0xasdf",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://example.com/metadata.json",
			},
			wantErr:     true,
			expectedErr: ErrInvalidEarningsReceiverAddress,
		},
		{
			name: "failed operator validation - wrong DelegationApproverAddress address",
			operator: Operator{
				Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				DelegationApproverAddress: "0x12",
				EarningsReceiverAddress:   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				StakerOptOutWindowBlocks:  100,
				MetadataUrl:               "https://example.com/metadata.json",
			},
			wantErr:     true,
			expectedErr: ErrInvalidDelegationApproverAddress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedErr, tt.operator.Validate(), "error should be equal")
		})
	}
}
