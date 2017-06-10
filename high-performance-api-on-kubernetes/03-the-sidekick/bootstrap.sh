#!/bin/bash

./create-configmap.sh $CONFIGMAP_REDIS_PORT \
                      $CONFIGMAP_REDIS_ADDR \
                      $COMPOSE_REDIS_PROXY_PORT \
                      $CONFIGMAP_MASTER_REDIS_PASS \
                      $CONFIGMAP_MASTER_REDIS_ADDR

sh ../01-the-tools/compose-ssl-tunnel/create-privkey-secret.sh

sh ../01-the-tools/compose-ssl-tunnel/create-ssl-tunnel-secrets.sh $COMPOSE_REDIS_ADDR $COMPOSE_REDIS_PORT $COMPOSE_REDIS_PROXY_IP $COMPOSE_REDIS_PROXY_PORT

kubectl apply -f ../01-the-tools/aggregator/kube.yaml

kubectl apply -f kube-redis-read.yaml && \
kubectl apply -f ../02-the-naive/updater/kube.yaml && \
cd ../02-the-naive && ./deploy-attacker.sh 1000 && cd -
