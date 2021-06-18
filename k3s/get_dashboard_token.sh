#!/bin/bash
sudo kubectl -n kubernetes-dashboard describe secret admin-user-token | sed -n -e "s/token:\s*\(\S*\)/\1/p"

