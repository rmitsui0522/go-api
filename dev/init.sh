#!/bin/bash -eu

source ./.env

docker-compose up -d --build
# dev/db_restore.sh

echo ""
echo "---------------------------------------------"
echo " [Sever start] API ENDPOINT:"
echo "---------------------------------------------"
echo "    Users:  http://localhost:$PORT/users      "
echo "    Health: http://localhost:$PORT/health     "
echo "---------------------------------------------"
echo ""
