#!/usr/bin/env bash

echo "Running podman-compose..."

podman-compose -f compose.yaml up -d  

podman ps -a
echo "Done"