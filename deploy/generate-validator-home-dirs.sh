#!/usr/bin/env bash

set -e
set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

VALIDATOR_IPS=$1

# convert validator ips variable from a string to an array
IFS=' ' read -r -a VALIDATOR_IPS <<< "${VALIDATOR_IPS}"

# create and initialize the the validator config dirs
mkdir -p deploy/node-config
rm -rf deploy/node-config/validator*
VALIDATOR_MNEMONICS=("gun quick banner word mutual pet sort run illness behind pull stock crazy talk actor icon help gym young census decorate swamp two plunge"
 "mule multiply combine frown aim window top weekend frown cancel turn token canoe thumb attitude flame execute purpose chest design winner enable coconut retire" "business bless fuel joy lady volcano odor tribe virus have effort rate mouse disease general view mention evoke lend expect frozen trend shrimp flavor")
for i in ${!VALIDATOR_IPS[@]}; do 
    MONIKER=validator${i}
    HOME_DIR=deploy/node-config/${MONIKER}
    echo ${VALIDATOR_MNEMONICS[$i]} | numid init --chain-id numi-testnet-1 --recover --home ${HOME_DIR} ${MONIKER}
    deploy/add-test-keys.sh ${HOME_DIR}
done

# update config files with external_address, persistent_peers and other values
P2P_PERSISTENT_PEERS=$(deploy/persistent-peers.sh "${VALIDATOR_IPS}")
for i in ${!VALIDATOR_IPS[@]}; do
  MONIKER=validator${i}
  HOME_DIR=deploy/node-config/${MONIKER}

  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.external_address" "tcp://${VALIDATOR_IPS[$i]}:26656"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.persistent_peers" "${P2P_PERSISTENT_PEERS}"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_cert_file" "/home/ubuntu/cert/fullchain.pem"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_key_file" "/home/ubuntu/cert/privkey.pem"
  dasel put bool -f ${HOME_DIR}/config/app.toml -p toml ".api.enable" true
  dasel put string -f ${HOME_DIR}/config/app.toml -p toml ".api.address" "tcp://localhost:1317"
done

# generate genesis transaction on each validator; for secondary validators, copy genesis transaction to primary validator
PRIMARY_HOME_DIR=deploy/node-config/validator0
for i in ${!VALIDATOR_IPS[@]}; do
  MONIKER=validator${i}
  HOME_DIR=deploy/node-config/${MONIKER}
  yes | numid keys delete ${MONIKER}-key --keyring-backend test --home ${HOME_DIR} 2>/dev/null || :
  echo ${VALIDATOR_MNEMONICS[$i]} | numid keys add ${MONIKER}-key --keyring-backend test --home ${HOME_DIR} --recover
  ADDR=$(numid keys show ${MONIKER}-key -a --keyring-backend test --home ${HOME_DIR})
  numid add-genesis-account --keyring-backend test --home ${HOME_DIR} ${ADDR} 2000000000unumi || :
  numid gentx ${MONIKER}-key 1000000000unumi --chain-id numi-testnet-1 --moniker=${MONIKER} --keyring-backend test --home ${HOME_DIR}

  if [[ ${i} != "0" ]]; then
    cp ${HOME_DIR}/config/gentx/* ${PRIMARY_HOME_DIR}/config/gentx/    
    numid add-genesis-account --keyring-backend test --home ${PRIMARY_HOME_DIR} ${ADDR} 2000000000unumi || :
  fi
done

# collect genesis transactions on primary validator node
numid collect-gentxs --home ${PRIMARY_HOME_DIR}

# add genesis accounts for test accounts
for KEY_NAME in oliver olivia alice bob carol; do
  ADDR=$(numid keys show ${KEY_NAME} -a --keyring-backend test --home ${PRIMARY_HOME_DIR})
  numid add-genesis-account --keyring-backend test --home ${PRIMARY_HOME_DIR} ${ADDR} 2000000000unumi || :
done

# update module parameters within genesis file on primary node
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.crisis.constant_fee.denom" "unumi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.gov.deposit_params.min_deposit.[0].denom" "unumi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.gov.voting_params.voting_period" "60s"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.mint.params.mint_denom" "unumi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.staking.params.bond_denom" "unumi"
ORACLE_OLIVER=$(numid keys show oliver -a --keyring-backend test --home ${PRIMARY_HOME_DIR})
ORACLE_OLIVIA=$(numid keys show olivia -a --keyring-backend test --home ${PRIMARY_HOME_DIR})
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.numi.params.identityVerifiers" "${ORACLE_OLIVER},${ORACLE_OLIVIA}"

# copy genesis file from primary node to secondary nodes
for i in ${!VALIDATOR_IPS[@]}; do
  if [[ "${i}" != "0" ]]; then
    MONIKER=validator${i}
    HOME_DIR=deploy/node-config/${MONIKER}
    cp ${PRIMARY_HOME_DIR}/config/genesis.json ${HOME_DIR}/config/
  fi
done
