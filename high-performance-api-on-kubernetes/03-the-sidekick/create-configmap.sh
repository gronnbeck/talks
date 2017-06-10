#!/bin/bash

kubectl delete configmap redis-sidekick
kubectl delete configmap redis-write

sh create-redis-config.sh $1 $2 $3 $4 && \
kubectl create configmap redis-sidekick --from-file=redis.conf

kubectl create configmap redis-write \
  --from-literal=host=$5 \
  --from-literal=pass=$4
