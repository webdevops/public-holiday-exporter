#!/bin/bash

# Config example:
#
# COUNTRY_TZ="DE,Europe/Berlin CH,Europe/Zurich"
# PROXY_ADDR="proxy:3128"
# ENDPOINT="https://my-endpoint:8080"

function add_config(){
    echo -e "$@" >> /tmp/config.yaml
}

add_config "countries:"
for i in $COUNTRY_TZ; do
    IFS=","
    key="- country"
    for j in $i; do
        add_config "$key: \"$j\""
        key="  timezone"
    done
    add_config
    IFS=" "
done

if [ -n "$PROXY_ADDR" ]; then
    add_config "proxy: \"$PROXY_ADDR\""
fi

if [ -n "$ENDPOINT" ]; then
    add_config "endpoint: \"$ENDPOINT\""
fi

mv /tmp/config.yaml /app/config.yaml

/app/public-holiday-exporter -c /app/config.yaml
