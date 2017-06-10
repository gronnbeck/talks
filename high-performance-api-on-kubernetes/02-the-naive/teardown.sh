#!/bin/bash

kubectl delete deployment vegetaserver
kubectl delete deployment aggregator
kubectl delete deployment redis-api
kubectl delete deployment updater
