#!/bin/bash

ATTACK_RATE=$1

cat kube-vegeta-api.yaml.tmpl | \
  sed -e s/'$ATTACK_RATE'/$ATTACK_RATE/g \
  | kubectl apply -f -
