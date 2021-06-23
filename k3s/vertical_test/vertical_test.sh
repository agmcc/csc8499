#!/bin/bash
# Example usage: vertical_test.sh 200M 1 2
MEMORY=$1
CPU=$2
DIFFICULTY=$3
echo "Memory: ${MEMORY}"
echo "CPU: ${CPU}"
echo "DIFFICULTY: ${DIFFICULTY}"
export MEMORY
export CPU
export DIFFICULTY
envsubst < vertical_test.tmpl > tmp.yml
sudo kubectl replace --force -f tmp.yml
rm tmp.yml
