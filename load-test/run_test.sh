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

host=http://instance-2:8080
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

