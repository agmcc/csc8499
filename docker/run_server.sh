#!/bin/bash
docker run \
--pull always \
--rm \
-e HOST=$(hostname) \
-e DIFFICULTY=4 \
-p "8080:8080" \
--cpus=4 \
--memory=200m \
ghcr.io/agmcc/csc8499/go-server
