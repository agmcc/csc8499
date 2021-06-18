#!/bin/bash
exposed=9000
sudo kubectl port-forward -n kubernetes-dashboard service/kubernetes-dashboard ${exposed}:443 --address 0.0.0.0 > /dev/null 2>&1 &

