package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

type DataRecord struct {
    Data string `json:"data"`
}

func (s *SmartContract) MigrateData(ctx contractapi.TransactionContextInterface, key string) error {
    dataAsBytes, err := ctx.GetStub().GetState("onchainA_" + key)
    if err != nil || dataAsBytes == nil {
        return fmt.Errorf("data not found in Onchain A")
    }

    err = ctx.GetStub().PutState("onchainB_" + key, dataAsBytes)
    if err != nil {
        return fmt.Errorf("failed to migrate data to Onchain B")
    }

    return nil
}

func (s *SmartContract) MigrateDataToOffChain(ctx contractapi.TransactionContextInterface, key string) error {
    dataAsBytes, err := ctx.GetStub().GetState("onchainA_" + key)
    if err != nil || dataAsBytes == nil {
        return fmt.Errorf("data not found in Onchain A")
    }

    var dataRecord DataRecord
    json.Unmarshal(dataAsBytes, &dataRecord)

    eventPayload := map[string]string{
        "key":  key,
        "data": dataRecord.Data,
    }

    payloadAsBytes, err := json.Marshal(eventPayload)
    if err != nil {
        return fmt.Errorf("failed to marshal event payload")
    }

    err = ctx.GetStub().SetEvent("DataCopiedToOffChain", payloadAsBytes)
    if err != nil {
        return fmt.Errorf("failed to emit off-chain event")
    }

    return nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        fmt.Printf("Error creating chaincode: %v", err)
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting chaincode: %v", err)
    }
}
