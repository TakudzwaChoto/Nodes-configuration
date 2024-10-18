#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

# imports
. scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER1_CA=${PWD}/organizations/ordererOrganizations/itblanket.org/orderer1/tlsca/tlsca-cert.pem
export ORDERER2_CA=${PWD}/organizations/ordererOrganizations/itblanket.org/orderer2/tlsca/tlsca-cert.pem
export ORDERER3_CA=${PWD}/organizations/ordererOrganizations/itblanket.org/orderer3/tlsca/tlsca-cert.pem
export PEER0_WATER_QUALITY_DATA , UPDATING AND RETRIEVALS,STORAGE_CA=${PWD}/organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/tlsca/tlsca.water_quality_data , updating and retrievals,storage-cert.pem
export PEER0_GOVERNMENT_CA=${PWD}/organizations/peerOrganizations/government/tlsca/tlsca.government-cert.pem
export PEER0_INCENTIVES_CONTRANTS_CA=${PWD}/organizations/peerOrganizations/incentives_contrants/tlsca/tlsca.incentives_contrants-cert.pem

# Set environment variables for the peer org
setGlobals() {
  PEER=$1
  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$2
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  infoln "Using organization ${USING_ORG}"
  if [ "$USING_ORG" == "water_quality_data , updating and retrievals,storage" ]; then
    export CORE_PEER_LOCALMSPID="Water_quality_data,UpdatingAndRetrievals,StorageMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_WATER_QUALITY_DATA , UPDATING AND RETRIEVALS,STORAGE_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/users/Admin@water_quality_data , updating and retrievals,storage/msp
    if [ $PEER -eq 0 ]; then
      export CORE_PEER_ADDRESS=localhost:4444
    elif [ $PEER -eq 1 ]; then
      export CORE_PEER_ADDRESS=localhost:4454
    elif [ $PEER -eq 2 ]; then
      export CORE_PEER_ADDRESS=localhost:4464
    fi
  elif [ "$USING_ORG" == "government" ]; then
    export CORE_PEER_LOCALMSPID="GovernmentMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_GOVERNMENT_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/government/users/Admin@government/msp
    if [ $PEER -eq 0 ]; then
      export CORE_PEER_ADDRESS=localhost:5555
    elif [ $PEER -eq 1 ]; then
      export CORE_PEER_ADDRESS=localhost:5565
    elif [ $PEER -eq 2 ]; then
      export CORE_PEER_ADDRESS=localhost:5575
    fi
  elif [ "$USING_ORG" == "incentives_contrants" ]; then
    export CORE_PEER_LOCALMSPID="Incentives_contrantsMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_INCENTIVES_CONTRANTS_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/incentives_contrants/users/Admin@incentives_contrants/msp
    if [ $PEER -eq 0 ]; then
      export CORE_PEER_ADDRESS=localhost:6666
    elif [ $PEER -eq 1 ]; then
      export CORE_PEER_ADDRESS=localhost:6676
    elif [ $PEER -eq 2 ]; then
      export CORE_PEER_ADDRESS=localhost:6686
    fi

  else
    errorln "ORG Unknown"
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# Set environment variables for use in the CLI container 
setGlobalsCLI() {
  setGlobals 0 $1

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  if [ "$USING_ORG" == "water_quality_data , updating and retrievals,storage" ]; then
    export CORE_PEER_ADDRESS=peer0.water_quality_data , updating and retrievals,storage:4444
  elif [ "$USING_ORG" == "government" ]; then
    export CORE_PEER_ADDRESS=peer0.government:5555
  elif [ "$USING_ORG" == "incentives_contrants" ]; then
    export CORE_PEER_ADDRESS=peer0.incentives_contrants:6666

  else
    errorln "ORG Unknown"
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""

  # Loop through the input parameters as an array of strings
  for PARAM in "$@"; do
    setGlobals 0 "$PARAM"
    PEER="peer0.$PARAM"
    ## Set peer addresses
    if [ -z "$PEERS" ]; then
      PEERS="$PEER"
    else
      PEERS="$PEERS $PEER"
    fi

    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)

    ## Set path to TLS certificate
    CA="PEER0_${PARAM^^}"_CA
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
  done

  # Remove leading space for output
  PEERS="$(echo -e "$PEERS" | sed -e 's/^[[:space:]]*//')"
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}

