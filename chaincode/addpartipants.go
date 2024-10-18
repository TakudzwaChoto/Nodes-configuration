package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const parPrefix = "PAR"

// SmartContract of this fabric sample
type SmartContract struct {
	contractapi.Contract
}

// Addparticipants issues 
func (s *SmartContract) Addparticipants(ctx contractapi.TransactionContextInterface, args string) Response {
	participant := &Partcipant{}
	err := JSONtoObject([]byte(args), participant)

	// Generate a unique 
	parKey, err := ctx.GetStub().CreateCompositeKey(parPrefix, []string{participant.PartcipantID})

	isEmpExists, err := s.ParticipantExists(ctx, parKey)
	if isEmpExists {
		return BuildResponse("DUPLICATE", fmt.Sprintf("Participant data record already exists in the blockchain"), nil)
	}
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to add new participant to the blockchain"), nil)
	}

	objEmpBytes, err := ObjecttoJSON(participant)
	err = ctx.GetStub().PutState(parKey, objEmpBytes)
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to add new participant to the blockchain"), nil)
	}
	return BuildResponse("SUCCESS", fmt.Sprintf("New participant record has been added to the blockchain successfully."), nil)

}

// ParticipantExists 
func (s *SmartContract) ParticipantExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	participantJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return participantJSON != nil, nil
}

// ReadParticipant 
func (s *SmartContract) ReadParticipant(ctx contractapi.TransactionContextInterface, args string) Response {
	participant := &Partcipant{}
	err := JSONtoObject([]byte(args), participant)

	patKey, err := ctx.GetStub().CreateCompositeKey(patPrefix, []string{participant.PartcipantID})
	participantBytes, err := ctx.GetStub().GetState(patKey)

	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
	}
	if participantBytes == nil {
		return BuildResponse("ERROR", fmt.Sprintf("The participant %s does not exist", participant.FirstName), nil)
	}
	return BuildResponse("SUCCESS", "", participantBytes)
}

// Search participant record, checking the string starts with the key...(Remove '^' from regex to search anywhere in the string)
func (s *SmartContract) QueryByPartialKey(ctx contractapi.TransactionContextInterface, args string) Response {
	key := &SearchKey{}
	err := JSONtoObject([]byte(args), key)

	if err != nil {
		fmt.Println("Error when marshall json:", err)
	}

	keyVal := key.Key

	// Get the query iterator for the given key
	queryString := fmt.Sprintf(`{
        "selector": {
            "$or": [
                {"fName": {"$regex": "(?i)^%s"}},
                {"lName": {"$regex": "(?i)^%s"}},
                {"mobile": {"$regex": "(?i)^%s"}},
                {"address": {"$regex": "(?i)^%s"}}
            ]
        }
    }`, keyVal, keyVal, keyVal, keyVal)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		fmt.Println("error:", err)
		return BuildResponse("ERROR", fmt.Sprintf("Error occurred when query the database"), nil)
	}
	defer resultsIterator.Close()

	// Iterate over the results and create an array of participantRecord objects
	var records []*Partcipant
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return BuildResponse("ERROR", fmt.Sprintf("Error occurred when iterate records"), nil)
		}

		var record Participant
		err = json.Unmarshal(queryResponse.Value, &record)
		if err != nil {
			return BuildResponse("ERROR", fmt.Sprintf("Error occurred when Unmarshal object"), nil)
		}
		records = append(records, &record)
	}

	participantBytes, err := ObjecttoJSON(records)
	return BuildResponse("SUCCESS", "", participantBytes)
}

// Updateparticipant updates an existing participant in the world state with provided parameters.
func (s *SmartContract) UpdateParticipant(ctx contractapi.TransactionContextInterface, args string) Response {
	participant := &Participant{}
	err := JSONtoObject([]byte(args), participant)

	parKey, err := ctx.GetStub().CreateCompositeKey(parPrefix, []string{participant.PartcipantID})

	objEmpBytes, err := ObjecttoJSON(participant)
	err = ctx.GetStub().PutState(parKey, objEmpBytes)
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to update participant record in the blockchain"), nil)
	}
	return BuildResponse("SUCCESS", fmt.Sprintf("Participant record has been updated in the blockchain successfully."), nil)
}

// GetAllParticipant returns all participant found in world state
func (s *SmartContract) GetAllParticipant(ctx contractapi.TransactionContextInterface) Response {
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(patPrefix, []string{})
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
	}
	defer resultsIterator.Close()

	var participant []*Participant
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
		}

		var participant Partcipant
		err = json.Unmarshal(queryResponse.Value, &participant)
		if err != nil {
			return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
		}
		participants = append(participant, &participant)
	}
	participantBytes, err := ObjecttoJSON(participant)
	return BuildResponse("SUCCESS", "", participantBytes)
}

// DeleteParticipant 
func (s *SmartContract) DeleteParticipant(ctx contractapi.TransactionContextInterface, args string) Response {
	participant := &Partcipant{}
	err := JSONtoObject([]byte(args), participant)

	parKey, err := ctx.GetStub().CreateCompositeKey(parPrefix, []string{participant.PartcipantID})

	participantBytes, err := ctx.GetStub().GetState(parKey)

	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to read participant data from the blockchain"), nil)
	}
	if participantBytes == nil {
		return BuildResponse("ERROR", fmt.Sprintf("The participant %s does not exist", participant.FirstName), nil)
	}

	err = ctx.GetStub().DelState(patKey)
	if err != nil {
		return BuildResponse("ERROR", fmt.Sprintf("Failed to delete participant record from the blockchain"), nil)
	}
	return BuildResponse("SUCCESS", "participant record deleted successfully.", nil)

}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create employee details chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting employee details chaincode: %s", err.Error())
	}
}
