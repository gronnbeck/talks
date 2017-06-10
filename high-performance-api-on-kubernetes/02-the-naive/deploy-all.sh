#!/bin/bash

kubectl apply -f naive/kube.yaml && \
kubectl apply -f updater/kube.yaml && \
./deploy-attacker.sh 1000

kubectl apply -f ../01-the-tools/aggregator/kube.yaml
