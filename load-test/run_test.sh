#!/bin/bash
source ./venv/bin/activate

if [ -z "$1" ]; then
    echo "Usage: "$0" TESTNAME"
    exit 0
fi

test=$1

if [ -d "$1" ]; then
    echo "Test directory '${1}' already exists, exiting"
    exit 1
fi

mkdir $test && cd $test

locust \
--locustfile ../locustfile.py \
--headless \
--host http://pilab-01:8080 \
--users 100 \
--spawn-rate 4 \
--run-time 90s \
--csv $test \
--html "${test}.html" \
--only-summary
