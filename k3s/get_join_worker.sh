#!/bin/bash
token=$(sudo cat /var/lib/rancher/k3s/server/node-token)
#command="curl -sfL https://get.k3s.io | K3S_URL=https://$(hostname):6443 K3S_TOKEN=${token} INSTALL_K3S_VERSION=v0.9.1 sh -"
command="curl -sfL https://get.k3s.io | K3S_URL=https://$(hostname):6443 K3S_TOKEN=${token} sh -"
echo $command
