#!/bin/bash
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 CPUS MEMORY DIFFICULTY"
    exit 0
fi

cpus=$1
memory=$2
difficulty=$3

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

