#!/usr/bin/env bash

podman ps -a
echo "Stoping all..."
podman stop $(podman ps -a -q)
echo "Removing all..."
podman rm $(podman ps -a -q)
podman ps -a
echo "Done"
