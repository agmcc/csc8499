#!/bin/bash
sudo kubectl delete ns kubernetes-dashboard
sudo kubectl -n kubernetes-dashboard delete serviceaccount admin-user
sudo kubectl -n kubernetes-dashboard delete clusterrolebinding admin-user

