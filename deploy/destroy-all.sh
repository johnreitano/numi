#!/usr/bin/env bash
# set -x
set -e

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

terraform -chdir=deploy destroy -auto-approve -var="num_validator_instances=0" -var="num_seed_instances=0" -var="create_explorer=false" -var="domain_prefix=testnet-" -var-file="dns.tfvars"
