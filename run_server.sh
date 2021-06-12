#!/bin/bash
docker run -e HOST=$(hostname) -p "8080:8080" ghcr.io/agmcc/csc8499/go-server:cd364ad6a08ca81b47d1c934abf64d9c53731cf7

