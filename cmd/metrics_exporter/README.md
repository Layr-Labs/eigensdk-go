# Metrics Exporter Sidecar

This binary is meant to be run as a sidecar, to periodically query the contracts and export the economic metrics on a prometheus endpoint.

## Usage

```bash
go run . --avs-name my-avs --quorum-names-dict '{"0":"eth","1":"mycoin"}' --operator-addr 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --registry-coordinator-addr 0x9E545E3C0baAB3E08CdfD552C960A1050f373042 --metrics-listening-addr localhost:9095
```