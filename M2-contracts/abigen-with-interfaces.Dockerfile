# use the below command to build this docker image locally
# docker build -t abigen-with-interfaces -f abigen-with-interfaces.Dockerfile .
FROM golang:1.21 as build

# this installs our fork of abigen with also creates interfaces for the binding structs
# See https://github.com/ethereum/go-ethereum/pull/28279
RUN apt update && \
    apt install git -y && \
    git clone https://github.com/samlaf/go-ethereum.git && \
    cd go-ethereum/cmd/abigen && \
    git checkout issue-28278/abigen-interfaces && \
    go build .


FROM debian:latest
COPY --from=build /go/go-ethereum/cmd/abigen/abigen /usr/local/bin/abigen
ENTRYPOINT [ "abigen"]