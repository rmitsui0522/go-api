#!/bin/bash -eu

source ./.env
DUMP_FILE="dev/sql/dump.sql"

if [ ! -e $DUMP_FILE ]; then
  mkdir dev/sql/
fi

docker exec -it $DB_CONTAINER_NAME sh -c "mysqldump -h localhost -p$DB_ROOT_PASSWORD --databases $DB_DATABASE" >$DUMP_FILE
