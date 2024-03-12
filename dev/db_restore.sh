#!/bin/bash -eu

source ./.env

if [ -e $DUMP_FILE ]; then
  docker exec -i $DB_CONTAINER_NAME sh -c "mysql -h localhost -p$DB_ROOT_PASSWORD" <$DUMP_FILE
fi
