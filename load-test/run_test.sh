#!/bin/bash
source ./venv/bin/activate

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 HOST TEST"
    exit 0
fi

host=$1
test=$2

if [ -d "$test" ]; then
    echo "Test directory '${test}' already exists, exiting"
    exit 1
fi

mkdir $test && cd $test

users=100
spawn=4
time=90s

locust \
--locustfile ../locustfile.py \
--headless \
--host $host \
--users $users \
--spawn-rate $spawn \
--run-time $time \
--csv $test \
--html "${test}.html" \
--only-summary

