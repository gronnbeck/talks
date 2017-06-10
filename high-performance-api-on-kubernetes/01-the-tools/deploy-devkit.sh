#!/bin/bash

kubectl apply -f helloworldserver/kube.yaml && \
kubectl apply -f vegetaserver/kube-helloworld.yaml && \
kubectl apply -f aggregator/kube.yaml
