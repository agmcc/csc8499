#!/bin/bash
# Run as root
sudo su

# Install
apt update
apt install -y strongswan

# Config
cp gcloud.conf /etc
conf="include /etc/gcloud.conf"
confFile=/etc/ipsec.conf
grep -qxF $conf $confFile || echo $conf >> $confFile

# Secret
cp gcloud.secrets /etc
read -p "Enter PSK: " psk
secretsFile=/etc/ipsec.secrets
sed -i "s/<PSK>/${PSK}/" $secretsFile
secret="include /etc/gcloud.secrets"
grep -qxF $secret $secretsFile || echo $secret >> $secretsFile

# Restart
service ipsec restart
