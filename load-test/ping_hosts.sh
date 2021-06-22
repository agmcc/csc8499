#!/bin/bash
count=20

hosts=(
    "DESKTOP-6SV4J7G" \
    "pilab-01" \
    "pilab-05" \
    "pilab-06" \
    "instance-1" \
    "instance-2" \
    "instance-3" \
)

hostname=$(hostname)

for h in ${hosts[@]}
do
    if [ $h != ${hostname} ]
    then
        echo "Pinging ${h}..."
        result=$(ping $h -c $count -q | grep rtt)
        echo "${hostname}:${h} ${result}"
    fi
done
