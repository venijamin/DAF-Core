#!/bin/bash

set -e

DB_CONTAINER=$(docker ps | grep 0.0.0.0:5432 | awk '{print $1}')
if [ -n "$DB_CONTAINER" ]; then
  docker stop "$DB_CONTAINER"
  docker rm "$DB_CONTAINER"
fi
sudo rm -rf ~/.daf/data
bash ./start-database.sh