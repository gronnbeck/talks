#! /bin/bash

set -e

_name="aggregator"

#SHA=$(git rev-parse --short HEAD)
#All builds live in production
CONTAINER_TAG="gronnbeck/$_name"

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -ldflags '-s -w' -installsuffix cgo -o output/$_name *.go

# Use an output folder containing only the Go binaries and the Dockerfile to minimize
# the Docker context
cp Dockerfile output
docker build -t $CONTAINER_TAG .
