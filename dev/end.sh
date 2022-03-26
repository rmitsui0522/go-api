#!/bin/bash -eu

./dev/db_dump.sh
docker-compose down
