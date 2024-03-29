#!/bin/bash -eu

source ./.env

if [ ! -e $DUMP_FILE ]; then
  mkdir dev/sql/
fi

docker exec -it $DB_CONTAINER_NAME sh -c "mysqldump -h localhost -p$MYSQL_ROOT_PASSWORD --databases $DB_DATABASE" >$DUMP_FILE
