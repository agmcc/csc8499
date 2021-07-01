#!/bin/bash
if [ -n "$1" ]
then
    namespace=$1
else
    namespace=default
fi

echo "Creating secret for namespace: ${namespace}"

read -p "Enter GitHub container registry PAT: " password

sudo kubectl create secret docker-registry ghcr \
--docker-server=ghcr.io \
--docker-username=agmcc \
--docker-password=$password \
-n $namespace
