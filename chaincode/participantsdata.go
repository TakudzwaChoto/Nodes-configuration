/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Define structs to be used by chaincode

type User struct {
	UserID    string `json:"userId,required"`
	Name      string `json:"name"`
	Password  string `json:"password,required"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	PaymentID string `json:"paymentID"`
	Timestamp string `json:"timeStamp"`
}

type Partcipant struct {
	PartcipantID    string `json:"participantID"`
	FirstName       string `json:"fName"`
	LastName        string `json:"lName"`
	DOB             string `json:"dob"`
	Gender          string `json:"gender"`
	Mobile          string `json:"mobile"`
	EmergencyNumber string `json:"emergency_phone"`
	Address         string `json:"address"`
}

type SearchKey struct {
	Key string `json:"searchString"`
}

type PartcipantPvtData struct {
    Water-supply         string `json:"Water-supply"`
	Water-discharge      string `json:"Water-disharge"`
	Water-treatment      string `json:"Water-treatment"`
	Water-trading        string `json:"Water-trading"`
	Water-level          string `json:"Water-level"``

}
