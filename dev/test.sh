#!/bin/bash -eu

docker exec -it go-api bash -c "go test -v /app/pkg/..."
