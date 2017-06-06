#!/bin/sh

./smarter --read-redis-addr=$REDIS_ADDR_READ \
        --write-redis-addr=$REDIS_ADDR_WRITE \
        --write-redis-pass=$REDIS_PASS_WRITE
