#!/bin/bash

set -e

DB_CONTAINER=$(docker ps | grep 0.0.0.0:5432 | awk '{print $1}')
docker stop "$DB_CONTAINER"
docker rm "$DB_CONTAINER"
sudo rm -rf ~/.daf/data
bash ./start-database.sh