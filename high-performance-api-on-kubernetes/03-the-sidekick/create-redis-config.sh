#!/bin/bash

_port=$1
_masterHost=$2
_masterPort=$3
_masterPass=$4


cat redis.conf.tmpl | \
sed -e s/'$REDIS_MASTER_HOST'/${_masterHost:-127.0.0.1}/g | \
sed -e s/'$REDIS_MASTER_PORT'/${_masterPort:-6379}/g | \
sed -e s/'$REDIS_MASTER_PASS'/$_masterPass/g | \
sed -e s/'$REDIS_PORT'/${_port:-6379}/g \
> redis.conf
