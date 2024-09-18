#!/bin/sh

# Replace environment variables in config.template.json and create config.json
envsubst < /usr/share/nginx/html/config.json > /usr/share/nginx/html/config.2.json
mv /usr/share/nginx/html/config.2.json /usr/share/nginx/html/config.json

# Execute the Docker CMD
exec "$@"
