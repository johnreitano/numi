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

# get validator ip addresses
VALIDATOR_IPS=($(terraform -chdir=deploy output -json validator_ips | jq -r 'join(" ")'))

# generate config dirs
${SCRIPT_DIR}/generate-config-dirs.sh ${VALIDATOR_IPS[*]}

for i in ${!VALIDATOR_IPS[@]}; do
  echo configuring server ${i} with IP ${VALIDATOR_IPS[${i}]}
  
  # upload files to server
  cp -r deploy/node-config/validator${i} deploy/upload/numi-home
  (cd deploy; tar czf upload.tgz upload)
  scp -i ${SSH_PRIVATE_KEY_PATH} -pr deploy/upload.tgz ubuntu@${VALIDATOR_IPS[${i}]}:
  
  # extract files and run configuration script
  RPC_HOST=${DOMAIN_PREFIX}validator-${i}-rpc.${DNS_ZONE_NAME}
  API_HOST=${DOMAIN_PREFIX}validator-${i}-api.${DNS_ZONE_NAME}

  ssh -i ${SSH_PRIVATE_KEY_PATH} ubuntu@${VALIDATOR_IPS[${i}]} "rm -rf upload && tar xzf upload.tgz && upload/configure-server.sh ${TLS_CERTIFICATE_EMAIL} ${RPC_HOST} ${API_HOST}"
done
