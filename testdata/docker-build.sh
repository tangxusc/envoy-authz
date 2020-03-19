#!/usr/bin/env bash
cd ../cmd/ && CGO_ENABLED=0 GOOS=linux go build -o main main.go
cd ../ && docker build . -t envoy-auth:v1
rm -rf cmd/main
cd testdata/
k apply -f deployment.yaml -f filter.yaml -f istio.yaml