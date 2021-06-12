#!/bin/bash
token=$(docker swarm join-token worker -q)
command="docker swarm join --token ${token} $(hostname):2377"
echo $command

