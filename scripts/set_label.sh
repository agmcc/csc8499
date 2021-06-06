#!/bin/bash
hostname=$1
label=$2
echo "Adding label '${label}' to ${hostname}"
node=$(docker node ls | grep $hostname | awk '{ print $1}')
docker node update --label-add $label $node
echo "Labels for node ${node} (${hostname}):"
docker node inspect $node | jq ".[] | .Spec.Labels"

