#!/bin/bash -eu

source ./.env

# ./dev/db_dump.sh
docker-compose down
docker-compose up -d --build
# dev/db_restore.sh

echo ""
echo "---------------------------------------------------------"
echo "    [Sever start] API ENDPOINT:"
echo "---------------------------------------------------------"
echo "    Health: http://localhost:$PORT/health"
echo "    Users:  http://localhost:$PORT/$API_BASE_URL/users"
echo "---------------------------------------------------------"
echo ""
