package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

type QualityEvent struct {
    Timestamp    int64  `json:"timestamp"`
    QualityLevel int64  `json:"qualityLevel"`
}

type QualityRecords struct {
    Recent    []QualityEvent `json:"recent"`
    Historical []QualityEvent `json:"historical"`
}

func (s *SmartContract) ReportQuality(ctx contractapi.TransactionContextInterface, user string, qualityLevel int64) error {
    currentTimestamp := time.Now().Unix()
    newEvent := QualityEvent{Timestamp: currentTimestamp, QualityLevel: qualityLevel}

    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil {
        return fmt.Errorf("failed to get records: %v", err)
    }

    var records QualityRecords
    if recordsAsBytes != nil {
        json.Unmarshal(recordsAsBytes, &records)
    }

    records.Recent = append(records.Recent, newEvent)

    recordsAsBytes, err = json.Marshal(records)
    if err != nil {
        return fmt.Errorf("failed to marshal records: %v", err)
    }

    return ctx.GetStub().PutState(user, recordsAsBytes)
}

func (s *SmartContract) GetRecentQualityRecords(ctx contractapi.TransactionContextInterface, user string) ([]QualityEvent, error) {
    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil {
        return nil, fmt.Errorf("failed to get records: %v", err)
    }

    var records QualityRecords
    if recordsAsBytes == nil {
        return nil, fmt.Errorf("no records found")
    }

    json.Unmarshal(recordsAsBytes, &records)

    return records.Recent, nil
}

func (s *SmartContract) MigrateToHistorical(ctx contractapi.TransactionContextInterface, user string, beforeTimestamp int64) error {
    recordsAsBytes, err := ctx.GetStub().GetState(user)
    if err != nil {
        return fmt.Errorf("failed to get records: %v", err)
    }

    var records QualityRecords
    if recordsAsBytes != nil {
        json.Unmarshal(recordsAsBytes, &records)
    }

    newRecent := []QualityEvent{}
    for _, event := range records.Recent {
        if event.Timestamp < beforeTimestamp {
            records.Historical = append(records.Historical, event)
        } else {
            newRecent = append(newRecent, event)
        }
    }

    records.Recent = newRecent

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
