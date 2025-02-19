#!/bin/bash

set -e

DB_DIR=~/.daf/data
DB_NAME=daf-db

echo "Checking directory..."
mkdir -p "$DB_DIR"

docker build -t $DB_NAME ../database
docker run -p 5432:5432 -v "$DB_DIR":/var/lib/postgresql/data $DB_NAME