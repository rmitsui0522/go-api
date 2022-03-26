#!/bin/bash -eu

docker-compose up -d --build

# .envファイル内の変数を参照する
source ./.env

echo ""
echo "---------------------------------------------"
echo " [Sever start] API ENDPOINT:"
echo "---------------------------------------------"
echo "    Users:  http://localhost:$PORT/users      "
echo "    Health: http://localhost:$PORT/health     "
echo "---------------------------------------------"
echo ""
