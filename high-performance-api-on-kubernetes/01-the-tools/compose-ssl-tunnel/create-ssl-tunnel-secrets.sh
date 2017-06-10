#!/bin/bash

kubectl create secret generic compose-ssl-tunnel \
  --from-literal=host=$1 \
  --from-literal=port=$2 \
  --from-literal=dst_ip=$3 \
  --from-literal=src_port=$4
