#!/usr/bin/env bash
k delete -f gateway.yaml
k delete -f deployment.yaml
k delete -f filter.yaml
sleep 10
docker rmi envoy-auth:v1