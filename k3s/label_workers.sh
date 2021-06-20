#!/bin/bash
for worker in $(sudo kubectl get nodes | grep "<none>" | awk '{print $1}')
do
  sudo kubectl label node ${worker} node-role.kubernetes.io/worker=''
done

