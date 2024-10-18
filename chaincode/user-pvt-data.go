package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const patPvtCollectionName = "patPvtCollection"

func (s *SmartContract) AddParticipantPvtData(ctx contractapi.TransactionContextInterface) Response {
	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Error getting transient"), nil)
	}

	// Private data is passed in transient field, instead of func args
	transientPvtJSON, ok := transientMap["participant_pvt_properties"]
	if !ok {
		return BuildResponse("ERROR", fmt.Sprintf("Private data not found in the transient map input"), nil)
	}

	var patPrivate ParticipantPvtData
	err = json.Unmarshal(transientPvtJSON, &patPrivate)

	pvtJSONasBytes, err := json.Marshal(patPrivate)
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to marshal private data into JSON"), nil)
	}

	// Create a composite key to save participant water records separately
	// clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	// patKey, err := ctx.GetStub().CreateCompositeKey(patPrivate.ParticipantID, []string{clientMSPID})

	// Uncomment these lines to save orgs private data which need not share with others

	// ownerCollection, err := getCollectionName(ctx)
	// if err != nil {
	// 	return BuildResponse("ERROR", fmt.Sprintf("Failed to get Private collection name from the config file"), nil)
	// }

	// Save private data collection
	err = ctx.GetStub().PutPrivateData(patPvtCollectionName, patPrivate.ParticipantID, pvtJSONasBytes)

	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to add participant private data to the blockchain. Please check your permission."), nil)
	}
	return BuildResponse("SUCCESS", fmt.Sprintf("Participant private data added to the blockchain successfully."), nil)

}

func (s *SmartContract) QueryParticipantPvtRecordById(ctx contractapi.TransactionContextInterface, args string) Response {
	pvtObj := &ParticipantPvtData{}
	err := JSONtoObject([]byte(args), pvtObj)
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to unmarshal the JSON object"), nil)
	}

	// ownerCollection, err := getCollectionName(ctx)
	// if err != nil {
	// 	return BuildResponse("ERROR", fmt.Sprintf("Failed to Private collection name from the config file"), nil)
	// }

	patPvtBytes, err := ctx.GetStub().GetPrivateData(patPvtCollectionName, pvtObj.ParticipantID)

	// No Asset found, return empty response
	if patPvtBytes == nil {
		return BuildResponse("ERROR", fmt.Sprintf("Water record does not exist or you do not have access to it."), nil)
	}
	return BuildResponse("SUCCESS", "", patPvtBytes)

	// To save Water quality records separately:

	// patPvtBytes, err := ctx.GetStub().GetPrivateDataByPartialCompositeKey(patPvtCollectionName, "ParticipantPvtData", []string{pvtObj.ParticipantID})

	// if err != nil {
	// 	return BuildResponse("ERROR", fmt.Sprintf("Failed to read water data from the blockchain"), nil)
	// }
	// defer patPvtBytes.Close()

	// var pvtDataset []*ParticipantPvtData
	// for patPvtBytes.HasNext() {
	// 	queryResponse, err := patPvtBytes.Next()
	// 	if err != nil {
	// 		return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
	// 	}

	// 	var pvtData ParticipantPvtData
	// 	err = json.Unmarshal(queryResponse.Value, &pvtData)
	// 	if err != nil {
	// 		return BuildResponse("ERROR", fmt.Sprintf("Failed to convert water data record to JSON"), nil)
	// 	}
	// 	pvtDataset = append(pvtDataset, &pvtData)
	// }
	// pvtDataBytes, err := ObjecttoJSON(pvtDataset)
	// return BuildResponse("SUCCESS", "", pvtDataBytes)
}
