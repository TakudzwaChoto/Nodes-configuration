# Water Quality Management - Hyperledger Fabric Network

This project implements a **Water Quality Management System** using Hyperledger Fabric to ensure secure and decentralized and private data sharing. The project structure includes the following key components:

## Project Structure

```bash
├── chaincode         # Smart contract logic for water quality management
├── client-ui         # User interface for interacting with the system
├── helia-server      # Backend server for network communication
├── network           # Hyperledger Fabric network setup and configuration
├── server-api        # API for handling client-server communication
```

# Water Quality Management - Hyperledger Fabric Network

 The network is built using the `fabric-samples` repository, with the following key files contributing to the setup and operation of the Fabric network.

## Key Files

- **`configtx.yaml`**: Defines channel configurations and policies for the network.
- **`crypto-config.yaml`**: Generates cryptographic materials (certificates and keys) for organizations, peers, and orderers.
- **`docker-compose-cli.yaml`**: Defines the CLI container to interact with the network (e.g., chaincode installation).
- **`docker-compose-couch.yaml`**: Sets up CouchDB as an alternative state database.
- **`docker-compose.yaml`**: Core Docker Compose file for starting the network (peers, orderers, CA, etc.).
- **`connection.yaml`**: Configures client applications for connecting to the network.
- **Chaincode files** (located in `fabric-samples/chaincode/`): Define the business logic (smart contracts) executed on the network.
- **`byfn.sh` / `eyfn.sh`**: Automates network setup and configuration.
- **`genesis.block`**: Genesis block file for starting the ordering service.
- **`mychannel.tx`**: Channel transaction file for creating a channel.
- **Anchor peer update files**: Configurations for updating anchor peers in the network.
- **`.env` (optional)**: Contains environment variables for custom network configuration.

## How to Start the Network

1. **Navigate to the network directory**:
   ```bash
   cd fabric-samples/test-network
   cd network
   ./network.sh up
   ./network.sh createChannel
   ./network.sh deployCC
 
### 1. Generate Channel Artifacts
Before starting the network, generate the necessary channel artifacts using the `configtxgen` tool.

```bash
# Generate the Genesis Block
configtxgen -profile <ProfileName> -channelID system-channel -outputBlock ./channel-artifacts/genesis.block

# Generate the Channel Creation Transaction
configtxgen -profile <ProfileName> -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID mychannel
```
You can Use Docker Compose to bring up the Hyperledger Fabric network, including peers, orderers, and the Certificate Authority (CA).

Copy code
# Navigate to the network directory
cd network

# Start the network
docker-compose -f docker-compose.yaml up -d

Once the network is up, create a channel and have peers join the channel
# From the CLI container, create the channel
peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/mychannel.tx

# Join peer to the channel
peer channel join -b mychannel.block

Install the chaincode (smart contract) on all peers and instantiate it on the channel.
# Install chaincode on Peer0
peer chaincode install -n water-quality -v 1.0 -p github.com/chaincode/water_quality

# Instantiate the chaincode on the channel
peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n water-quality -v 1.0 -c '{"Args":["init"]}'

Invoke the chaincode to execute a transaction on the channel (e.g., updating water quality data)
# Invoke the chaincode function
peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n water-quality -c '{"Args":["updateWaterQuality","river1","75"]}'

Query the chaincode to retrieve water quality data.
# Query the water quality data for river1
peer chaincode query -C mychannel -n water-quality -c '{"Args":["queryWaterQuality","river1"]}'





