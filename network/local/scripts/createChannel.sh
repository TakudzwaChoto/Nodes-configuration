#!/bin/bash

# imports
. scripts/envVar.sh
. scripts/utils.sh

CHANNEL_NAME="$1"
DELAY="$2"
MAX_RETRY="$3"
VERBOSE="$4"
BFT="$5"
: ${CHANNEL_NAME:="gershom"}
: ${DELAY:="3"}
: ${MAX_RETRY:="5"}
: ${VERBOSE:="false"}
: ${BFT:=0}

: ${CONTAINER_CLI:="docker"}
: ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI}-compose"}
infoln "Using ${CONTAINER_CLI} and ${CONTAINER_CLI_COMPOSE}"

if [ ! -d "channel-artifacts" ]; then
	mkdir channel-artifacts
fi

createChannelGenesisBlock() {
	setGlobals 0 water_quality_data , updating and retrievals,storage
	which configtxgen
	if [ "$?" -ne 0 ]; then
		fatalln "configtxgen tool not found."
	fi
	local bft_true=$1
	set -x

	if [ $bft_true -eq 1 ]; then
		configtxgen -profile ChannelUsingBFT -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
	else
		configtxgen -profile ChannelUsingRaft -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
	fi
	res=$?
	{ set +x; } 2>/dev/null
	verifyResult $res "Failed to generate channel configuration transaction..."
}

createChannel() {
	# Poll in case the raft leader is not set yet
	local rc=1
	local COUNTER=1
	local bft_true=$1
	infoln "Adding orderers"
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ]; do
		sleep $DELAY
		set -x
			. scripts/orderer1.sh ${CHANNEL_NAME}> /dev/null 2>&1
			. scripts/orderer2.sh ${CHANNEL_NAME}> /dev/null 2>&1
			. scripts/orderer3.sh ${CHANNEL_NAME}> /dev/null 2>&1

		res=$?
		{ set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "Channel creation failed"
}

# joinChannel ORG
joinChannel() {
	FABRIC_CFG_PATH=$PWD/config/
	PEER=$1
	ORG=$2
	setGlobals $PEER $ORG
	local rc=1
	local COUNTER=1
	## Sometimes Join takes time, hence retry
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ]; do
		sleep $DELAY
		set -x
		peer channel join -b $BLOCKFILE >&log.txt
		res=$?
		{ set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "After $MAX_RETRY attempts, peer${PEER}.${ORG} has failed to join channel '$CHANNEL_NAME' "
}

setAnchorPeer() {
	ORG=$1
	${CONTAINER_CLI} exec cli ./scripts/setAnchorPeer.sh $ORG $CHANNEL_NAME
}

# Create channel genesis block
FABRIC_CFG_PATH=$PWD/config/
BLOCKFILE="./channel-artifacts/${CHANNEL_NAME}.block"

infoln "Generating channel genesis block '${CHANNEL_NAME}.block'"
FABRIC_CFG_PATH=${PWD}/configtx
if [ $BFT -eq 1 ]; then
	FABRIC_CFG_PATH=${PWD}/bft-config
fi
createChannelGenesisBlock $BFT

## Create channel
infoln "Creating channel ${CHANNEL_NAME}"
createChannel $BFT
successln "Channel '$CHANNEL_NAME' created"

## Join all the peers to the channel
infoln "Joining water_quality_data , updating and retrievals,storage peer0 to the channel..."
joinChannel 0 water_quality_data , updating and retrievals,storage
infoln "Joining water_quality_data , updating and retrievals,storage peer1 to the channel..."
joinChannel 1 water_quality_data , updating and retrievals,storage
infoln "Joining water_quality_data , updating and retrievals,storage peer2 to the channel..."
joinChannel 2 water_quality_data , updating and retrievals,storage
infoln "Joining government peer0 to the channel..."
joinChannel 0 government
infoln "Joining government peer1 to the channel..."
joinChannel 1 government
infoln "Joining government peer2 to the channel..."
joinChannel 2 government
infoln "Joining incentives_contrants peer0 to the channel..."
joinChannel 0 incentives_contrants
infoln "Joining incentives_contrants peer1 to the channel..."
joinChannel 1 incentives_contrants
infoln "Joining incentives_contrants peer2 to the channel..."
joinChannel 2 incentives_contrants

## Set the anchor peers for each org in the channel
infoln "Setting anchor peer for water_quality_data , updating and retrievals,storage..."
setAnchorPeer water_quality_data , updating and retrievals,storage
infoln "Setting anchor peer for government..."
setAnchorPeer government
infoln "Setting anchor peer for incentives_contrants..."
setAnchorPeer incentives_contrants


## Set the anchor peers for each org in the channel

successln "Channel '$CHANNEL_NAME' joined"

