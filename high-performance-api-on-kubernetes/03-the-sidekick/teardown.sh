#!/bin/bash

kubectl delete configmap redis-sidekick
kubectl delete configmap redis-write
kubectl delete secret compose-ssl-tunnel
kubectl delete secret ssl-tunnel-priv-key

kubectl delete deployment vegetaserver
kubectl delete deployment aggregator
kubectl delete deployment redis-api
kubectl delete deployment updater
