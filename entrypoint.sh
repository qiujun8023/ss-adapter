#!/bin/bash

APP_MANAGER_SERVER=${APP_MANAGER_SERVER:-'localhost'}
APP_MANAGER_PORT=${APP_MANAGER_PORT:-'8839'}
APP_API_URL=${APP_API_URL:-'http://ss.example.com'}
APP_NODE_ID=${APP_NODE_ID:-'1'}
APP_NODE_TOKEN=${APP_NODE_TOKEN:-'token'}
APP_SYNC_INTERVAL=${APP_SYNC_INTERVAL:-'30'}

jq -n -M \
--arg manager_server "$APP_MANAGER_SERVER" \
--arg manager_port "$APP_MANAGER_PORT" \
--arg api_url "$APP_API_URL" \
--arg node_id "$APP_NODE_ID" \
--arg node_token "$APP_NODE_TOKEN" \
--arg sync_interval "$APP_SYNC_INTERVAL" \
'{
  "manager_server": $manager_server,
  "manager_port": $manager_port | tonumber,
  "api_url": $api_url,
  "node_id": $node_id,
  "node_token": $node_token,
  "sync_interval": $sync_interval | tonumber
}' > config.json;

./ss-adapter -c config.json;