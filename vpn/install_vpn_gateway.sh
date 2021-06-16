#!/bin/bash
sudo apt update && sudo apt upgrade -y
sudo apt install strongswan -y

sudo cp /etc/sysctl.conf /etc/sysctl.conf.bkp
sudo su
cat >> /etc/sysctl.conf << EOF
net.ipv4.ip_forward = 1
net.ipv4.conf.all.accept_redirects = 0
net.ipv4.conf.all.send_redirects = 0
EOF
exit
sudo sysctl -p /etc/sysctl.conf

source vpn_config.sh
read -p "Enter PSK: " PSK
export PSK
envsubst < ipsec.conf.tmpl > ipsec.conf.tmp
envsubst < ipsec.secrets.tmpl > ipsec.secrets.tmp
sudo cp /etc/ipsec.conf /etc/ipsec.conf.bkp
sudo cp /etc/ipsec.secrets /etc/ipsec.secrets.bkp
sudo mv ipsec.conf.tmp /etc/ipsec.conf
sudo mv ipsec.secrets.tmp /etc/ipsec.secrets

sudo ipsec restart
sudo ipsec status
