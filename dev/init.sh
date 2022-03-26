#!/bin/bash -eu

docker-compose up -d --build

echo ""
echo "---------------------------------------------"
echo " [Sever start] Access URL:"
echo "---------------------------------------------"
echo "        UI: http://localhost:3000/           "
echo "---------------------------------------------"
echo ""
