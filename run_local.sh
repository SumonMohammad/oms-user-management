#!/bin/bash

echo "Starting server"

export OMS_USER_MANAGEMENT_CONSUL_PATH="oms-user-management"
export OMS_USER_MANAGEMENT_CONSUL_URL="localhost:8500"

# Uploading the configuration to Consul
curl --request PUT --data-binary @config/config.dev.yaml http://127.0.0.1:8500/v1/kv/oms-user-management

# Running the Go server with any passed arguments
go run cmd/*.go "$1"

