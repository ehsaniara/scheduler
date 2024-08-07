#!/bin/bash

# Check if backend servers are healthy and update NGINX configuration

BACKENDS=("delay-box-worker-ts-1:8088" "delay-box-worker-ts-2:8088" "delay-box-worker-ts-3:8088")
CONFIG_FILE="/etc/nginx/conf.d/upstream.conf"
TMP_FILE="/tmp/upstream.conf.tmp"

echo "upstream worker-ts {" > $TMP_FILE
for BACKEND in "${BACKENDS[@]}"; do
    if curl -s --head http://$BACKEND/ping | grep "200 OK" > /dev/null; then
        echo "    server $BACKEND;" >> $TMP_FILE
    else
        echo "    # server $BACKEND (down);" >> $TMP_FILE
    fi
done
echo "}" >> $TMP_FILE

mv $TMP_FILE $CONFIG_FILE
nginx -s reload
