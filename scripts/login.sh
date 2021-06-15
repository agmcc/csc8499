#!/bin/bash
read -p "Enter PAT: " pat
docker login ghcr.io --username agmcc --password $pat

