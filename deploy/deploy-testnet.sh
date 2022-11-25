#!/usr/bin/env bash

set -e
set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

TLS_CERTIFICATE_EMAIL="jreitano@gmail.com"
DOMAIN_PREFIX="testnet-"
DNS_ZONE_NAME="numi.oktryme.com"
SSH_PRIVATE_KEY_PATH="~/.ssh/id_rsa"

# build m1 mac client
make build-mac

# get ip addresses
VALIDATOR_IPS=($(terraform -chdir=deploy output -json validator_ips | jq -r 'join(" ")'))
SEED_IPS=($(terraform -chdir=deploy output -json seed_ips | jq -r 'join(" ")'))

# shut down all validator and seed nodes
for i in ${!VALIDATOR_IPS[@]}; do
  ssh -i ${SSH_PRIVATE_KEY_PATH} ubuntu@${VALIDATOR_IPS[${i}]} "sudo systemctl stop numi.service || :"
done  
for i in ${!SEED_IPS[@]}; do
  ssh -i ${SSH_PRIVATE_KEY_PATH} ubuntu@${SEED_IPS[${i}]} "sudo systemctl stop numi.service || :"
done

# generate validator config dirs
${SCRIPT_DIR}/generate-validator-home-dirs.sh "${VALIDATOR_IPS[*]}"

for i in ${!VALIDATOR_IPS[@]}; do
  echo configuring server ${i} with IP ${VALIDATOR_IPS[${i}]}
  
  # upload files to server
  rm -rf deploy/upload/numi-home
  cp -r deploy/node-config/validator${i} deploy/upload/numi-home
  (cd deploy; tar czf upload.tgz upload)
  scp -i ${SSH_PRIVATE_KEY_PATH} -pr deploy/upload.tgz ubuntu@${VALIDATOR_IPS[${i}]}:
  
  # extract files and run configuration script
  RPC_DOMAIN=${DOMAIN_PREFIX}validator-${i}-rpc.${DNS_ZONE_NAME}
  API_DOMAIN=${DOMAIN_PREFIX}validator-${i}-api.${DNS_ZONE_NAME}
  API_PORT=1317
  ssh -i ${SSH_PRIVATE_KEY_PATH} ubuntu@${VALIDATOR_IPS[${i}]} "rm -rf upload && tar xzf upload.tgz && upload/configure-server.sh ${TLS_CERTIFICATE_EMAIL} ${RPC_DOMAIN} ${API_DOMAIN} ${API_PORT}"
done

# generate seed config dirs
${SCRIPT_DIR}/generate-seed-home-dirs.sh "${VALIDATOR_IPS[*]}" "${SEED_IPS[*]}"

for i in ${!SEED_IPS[@]}; do
  echo configuring server ${i} with IP ${SEED_IPS[${i}]}
  
  # upload files to server
  rm -rf deploy/upload/numi-home
  cp -r deploy/node-config/seed${i} deploy/upload/numi-home
  (cd deploy; tar czf upload.tgz upload)
  scp -i ${SSH_PRIVATE_KEY_PATH} -pr deploy/upload.tgz ubuntu@${SEED_IPS[${i}]}:
  
  # extract files and run configuration script
  RPC_DOMAIN=${DOMAIN_PREFIX}seed-${i}-rpc.${DNS_ZONE_NAME}
  API_DOMAIN=${DOMAIN_PREFIX}seed-${i}-api.${DNS_ZONE_NAME}
  API_PORT=1317
  ssh -i ${SSH_PRIVATE_KEY_PATH} ubuntu@${SEED_IPS[${i}]} "rm -rf upload && tar xzf upload.tgz && upload/configure-server.sh ${TLS_CERTIFICATE_EMAIL} ${RPC_DOMAIN} ${API_DOMAIN} ${API_PORT}"
done
