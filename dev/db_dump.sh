#!/bin/bash -eu

source ./.env

docker exec -it $DB_CONTAINER_NAME sh -c "mysqldump -h localhost -p$DB_ROOT_PASSWORD -x $DB_DATABASE" >build/sql/dump.sql
