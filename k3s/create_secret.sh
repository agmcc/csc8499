#!/bin/bash
# Needed for private registry
read -p "Enter GitHub container registry PAT: " password

sudo kubectl create secret docker-registry ghcr \
--docker-server=ghcr.io \
--docker-username=agmcc \
--docker-password=$password
