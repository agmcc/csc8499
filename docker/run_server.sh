#!/bin/bash
if [ -z "$1" ]; then
    echo "Usage: $0 DIFFICULTY"
    exit 0
fi

difficulty=$1
cpus=2
memory=200m

echo "C ${cpus} M ${memory} D ${difficulty}"

docker run \
--pull always \
--rm \
-e HOST=$(hostname) \
-e DIFFICULTY=$difficulty \
-p "8080:8080" \
--cpus=$cpus \
--memory=$memory \
ghcr.io/agmcc/csc8499/go-server

