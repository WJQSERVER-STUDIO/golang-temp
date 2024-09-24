#!/bin/bash

if [ ! -f /data/caddy/config/Caddyfile ]; then
    cp /data/caddy/Caddyfile /data/caddy/config/Caddyfile
fi

if [ ! -f /data/go/config/config.yaml ]; then
    cp /data/go/config.yaml /data/go/config/config.yaml
fi

/data/caddy/caddy run --config /data/caddy/config/Caddyfile > /data/go/log/caddy.log 2>&1 &

/data/go/go > /data/ghproxy/log/run.log 2>&1 &

while [[ true ]]; do
    sleep 1
done    

