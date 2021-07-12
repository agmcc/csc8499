#!/bin/bash
activate=./venv/bin/activate

if [ ! -f $activate ]
then
    echo "Virtual env script not found: ${activate}"
    exit 1
fi

source $activate

if [ "$#" -ne 2 ]
then
    echo "Usage: $0 HOST TESTDIR"
    exit 0
fi

host=$1
testDir=$2

if [ -d "$testDir" ]
then
    echo "Test directory '${testDir}' already exists, exiting"
    exit 1
fi

test=$(basename $testDir)
locustFile=$(realpath locustfile.py)

if [ ! -f $locustFile ]
then
    echo "Locust file not found: ${locustFile}"
    exit 1
fi

mkdir -p $testDir && cd $testDir

users=100
spawn=7
time=60s

locust \
--locustfile $locustFile \
--headless \
--host $host \
--users $users \
--spawn-rate $spawn \
--run-time $time \
--csv $test \
--html "${test}.html" \
--only-summary
