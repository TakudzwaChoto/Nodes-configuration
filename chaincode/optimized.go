package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

type QualityEvent struct {
    Timestamp    uint64 `json:"timestamp"`
    QualityLevel int64  `json:"qualityLevel"`
}

type QualityRecords struct {
    Recent    map[uint64]int64 `json:"recent"`
    Historical []uint64        `json:"historical"`
}

func (s *SmartContract) ReportQuality(ctx contractapi.TransactionContextInterface, user string, qualityLevel int64) error {
    currentTimestamp := ctx.GetStub().GetTxTimestamp().Seconds
    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil {
        return fmt.Errorf("failed to get records: %v", err)
    }

    var records QualityRecords
    if recordsAsBytes != nil {
        json.Unmarshal(recordsAsBytes, &records)
    } else {
        records.Recent = make(map[uint64]int64)
    }

    records.Recent[currentTimestamp] = qualityLevel
    records.Historical = append(records.Historical, currentTimestamp)

    recordsAsBytes, err = json.Marshal(records)
    if err != nil {
        return fmt.Errorf("failed to marshal records: %v", err)
    }

    return ctx.GetStub().PutState(user, recordsAsBytes)
}

func (s *SmartContract) GetRecentQualityRecords(ctx contractapi.TransactionContextInterface, user string) ([]uint64, []int64, error) {
    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil || recordsAsBytes == nil {
        return nil, nil, fmt.Errorf("no records found")
    }

    var records QualityRecords
    json.Unmarshal(recordsAsBytes, &records)

    qualityLevels := make([]int64, len(records.Historical))
    for i, timestamp := range records.Historical {
        qualityLevels[i] = records.Recent[timestamp]
    }

    return records.Historical, qualityLevels, nil
}

func (s *SmartContract) MigrateToHistorical(ctx contractapi.TransactionContextInterface, user string, beforeTimestamp uint64) error {
    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil || recordsAsBytes == nil {
        return fmt.Errorf("no records found")
    }

    var records QualityRecords
    json.Unmarshal(recordsAsBytes, &records)

    for _, timestamp := range records.Historical {
        if timestamp < beforeTimestamp {
            delete(records.Recent, timestamp)
        }
    }

    recordsAsBytes, err = json.Marshal(records)
    if err != nil {
        return fmt.Errorf("failed to marshal records: %v", err)
    }

    return ctx.GetStub().PutState(user, recordsAsBytes)
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

