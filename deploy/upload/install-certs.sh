#!/usr/bin/env bash

set -x
set -e

TLS_CERTIFICATE_EMAIL=$1
RPC_DOMAIN=$2
API_DOMAIN=$3
API_PORT=$4

# kill nginx if it's running since it interferes with certbot
sudo systemctl stop nginx || :
sudo pkill nginx || :

# obtain cert for use by numi rpc service
sudo certbot certonly --standalone -d ${RPC_DOMAIN} -n --agree-tos --email ${TLS_CERTIFICATE_EMAIL}
mkdir -p ~/cert
sudo cp -f /etc/letsencrypt/live/${RPC_DOMAIN}/fullchain.pem ~/cert/
sudo cp -f /etc/letsencrypt/live/${RPC_DOMAIN}/privkey.pem ~/cert/
sudo chown -R ubuntu:ubuntu ~/cert/*.pem

# obtain cert for use by ngingx acting as reverse proxy for numi api service
sudo DEBIAN_FRONTEND=noninteractive apt install -y nginx
cat ~/upload/default.nginx | sed 's/__DOMAIN__/'${API_DOMAIN}'/g' | sed 's/__PORT__/'${API_PORT}'/g' >/tmp/default.nginx
sudo mv -f /tmp/default.nginx /etc/nginx/sites-enabled/default
sudo certbot --nginx -d ${API_DOMAIN} -n --agree-tos --email ${TLS_CERTIFICATE_EMAIL} --post-hook "certbot with nginx completed" && (sudo nginx -s stop || :)

sudo systemctl start nginx
