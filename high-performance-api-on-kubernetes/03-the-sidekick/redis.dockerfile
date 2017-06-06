FROM redis

# Run custom redis config
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
