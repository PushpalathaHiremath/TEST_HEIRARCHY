/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/ibm/ciav"
	"github.com/op/go-logging"
	"strings"
)

var myLogger = logging.MustGetLogger("customer_details")
var dummyValue = "99999"

type ServicesChaincode struct {
}

type Identification struct {
	CustomerId     string
	IdentityNumber string
	PoiType        string
	PoiDoc         string
	Source         string
}

type PersonalDetails struct {
	CustomerId   string
	FirstName    string
	LastName     string
	Sex          string
	EmailId      string
	Dob          string
	PhoneNumber  string
	Occupation   string
	AnnualIncome string
	IncomeSource string
	Source       string
}

type Kyc struct {
	CustomerId  string
	KycStatus   string
	LastUpdated string
	Source      string
}

type Address struct {
	CustomerId  string
	AddressId   string
	AddressType string
	DoorNumber  string
	Street      string
	Locality    string
	City        string
	State       string
	Pincode     string
	PoaType     string
	PoaDoc      string
	Source      string
}

type Customer struct {
	Identification  []Identification
	PersonalDetails PersonalDetails
	Kyc             Kyc
	Address         []Address
}

/*
   Deploy KYC data model
*/
func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	ciav.Initialize()
	ciav.CreateIdentificationTable(stub, args)
	ciav.CreateCustomerTable(stub, args)
	ciav.CreateKycTable(stub, args)
	ciav.CreateAddressTable(stub, args)
	return nil, nil
}

/*
  Add Customer record
*/
func (t *ServicesChaincode) addCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	var Cust Customer
	err := json.Unmarshal([]byte(string(args[0])), &Cust)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	for i := range Cust.Identification {
		ciav.AddIdentification(stub, []string{Cust.Identification[i].CustomerId, Cust.Identification[i].IdentityNumber, Cust.Identification[i].PoiType, Cust.Identification[i].PoiDoc,
			Cust.Identification[i].Source})
	}
	ciav.AddCustomer(stub, []string{Cust.PersonalDetails.CustomerId, Cust.PersonalDetails.FirstName, Cust.PersonalDetails.LastName,
		Cust.PersonalDetails.Sex, Cust.PersonalDetails.EmailId, Cust.PersonalDetails.Dob, Cust.PersonalDetails.PhoneNumber, Cust.PersonalDetails.Occupation,
		Cust.PersonalDetails.AnnualIncome, Cust.PersonalDetails.IncomeSource, Cust.PersonalDetails.Source})
	ciav.AddKYC(stub, []string{Cust.Kyc.CustomerId, Cust.Kyc.KycStatus, Cust.Kyc.LastUpdated, Cust.Kyc.Source})
	for i := range Cust.Address {
		ciav.AddAddress(stub, []string{Cust.Address[i].CustomerId, Cust.Address[i].AddressId, Cust.Address[i].AddressType,
			Cust.Address[i].DoorNumber, Cust.Address[i].Street, Cust.Address[i].Locality, Cust.Address[i].City, Cust.Address[i].State,
			Cust.Address[i].Pincode, Cust.Address[i].PoaType, Cust.Address[i].PoaDoc, Cust.Address[i].Source})
	}
	return nil, nil
}

/*
 Update customer record
*/
func (t *ServicesChaincode) updateCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	var Cust Customer
	err := json.Unmarshal([]byte(string(args[0])), &Cust)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	for i := range Cust.Identification {
		ciav.UpdateIdentification(stub, []string{Cust.Identification[i].CustomerId, Cust.Identification[i].IdentityNumber, Cust.Identification[i].PoiType, Cust.Identification[i].PoiDoc,
			Cust.Identification[i].Source})
	}
	ciav.UpdateCustomer(stub, []string{Cust.PersonalDetails.CustomerId, Cust.PersonalDetails.FirstName, Cust.PersonalDetails.LastName,
		Cust.PersonalDetails.Sex, Cust.PersonalDetails.EmailId, Cust.PersonalDetails.Dob, Cust.PersonalDetails.PhoneNumber, Cust.PersonalDetails.Occupation,
		Cust.PersonalDetails.AnnualIncome, Cust.PersonalDetails.IncomeSource, Cust.PersonalDetails.Source})
	ciav.UpdateKYC(stub, []string{Cust.Kyc.CustomerId, Cust.Kyc.KycStatus, Cust.Kyc.LastUpdated, Cust.Kyc.Source})
	for i := range Cust.Address {
		ciav.UpdateAddress(stub, []string{Cust.Address[i].CustomerId, Cust.Address[i].AddressId, Cust.Address[i].AddressType,
			Cust.Address[i].DoorNumber, Cust.Address[i].Street, Cust.Address[i].Locality, Cust.Address[i].City, Cust.Address[i].State,
			Cust.Address[i].Pincode, Cust.Address[i].PoaType, Cust.Address[i].PoaDoc, Cust.Address[i].Source})
	}
	return nil, nil
}

/*
   Invoke : addCIAV and updateCIAV
*/
func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	if function == "addCIAV" {
		// add customer
		return t.addCIAV(stub, args)
	} else {
		// update customer
		return t.updateCIAV(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

/*
	Get Customer record by customer id or PAN number
*/
func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "getCIAV" {
		return t.getCIAV(stub, args)
	} else if function == "getKYCStats" {
		return t.GetKYCStats(stub)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t *ServicesChaincode) getCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var jsonResp string
	var customerIds []string
	var err error

	var identificationStr string
	var customerStr string
	var kycStr string
	var addressStr string
	if args[0] == "PAN" {
		customerIds, err = ciav.GetCustomerID(stub, args[1])
		// jsonResp = "["
		for i := range customerIds {
			customerId := customerIds[i]
			identificationStr, err = ciav.GetIdentification(stub, customerId)
			customerStr, err = ciav.GetCustomer(stub, customerId)
			kycStr, err = ciav.GetKYC(stub, customerId)
			addressStr, err = ciav.GetAddress(stub, customerId)

			if i != 0 {
				jsonResp = jsonResp + ","
			}
			jsonResp = jsonResp + "{\"Identification\":" + identificationStr +
				",\"PersonalDetails\":" + customerStr +
				",\"KYC\":" + kycStr +
				",\"address\":" + addressStr + "}"
		}
		// jsonResp = jsonResp + "]"
	} else if args[0] == "CUST_ID" {
		customerId := args[1]
		identificationStr, err = ciav.GetIdentification(stub, customerId)
		customerStr, err = ciav.GetCustomer(stub, customerId)
		kycStr, err = ciav.GetKYC(stub, customerId)
		addressStr, err = ciav.GetAddress(stub, customerId)

		jsonResp = "{\"Identification\":" + identificationStr +
			",\"PersonalDetails\":" + customerStr +
			",\"KYC\":" + kycStr +
			",\"address\":" + addressStr + "}"
	} else {
		return nil, errors.New("Invalid arguments. Please query by CUST_ID or PAN")
	}
	callerRole, _ := stub.ReadCertAttribute("role")
	visibility := ciav.Helpdesk
	if string(callerRole) == "Superadmin" {
		visibility = ciav.Superadmin
	} else if string(callerRole) == "RelationalManager" {
		visibility = ciav.RelationalManager
	} else if string(callerRole) == "Manager" {
		visibility = ciav.Manager
	}

	var visibilityBuffer bytes.Buffer
	visibilityBuffer.WriteString("{")

	i := 0
	for key, value := range visibility {
		if i > 0 {
			visibilityBuffer.WriteString(",")
		}
		visibilityBuffer.WriteString("\"" + key + "\":\"" + value + "\"")
		i++
	}
	visibilityBuffer.WriteString("}")
	responseStr := "{\"data\":" + jsonResp +
		"\"visibility\":" + visibilityBuffer +
		"}"
	bytes, err := json.Marshal(responseStr)
	if err != nil {
		return nil, errors.New("Error converting kyc record")
	}
	return bytes, nil
}

func (t *ServicesChaincode) GetKYCStats(stub *shim.ChaincodeStub) ([]byte, error) {
	var err error

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: dummyValue}}
	columns = append(columns, col1)
	rows, err := GetAllRows(stub, "KYC", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retriving KYC details [%s]", err)
	}

	var kycBuffer bytes.Buffer
	// var compliantBuffer bytes.Buffer
	// var noncompliantBuffer bytes.Buffer
	var compliantCustomersCount int
	var noncompliantCustomersCount int
	var totalCustomers int

	for i := range rows {
		row := rows[i]
		totalCustomers++
		if row.Columns[2].GetString_() == "compliant" {
			compliantCustomersCount++
			// if compliantBuffer.String() != "" {
			// 	compliantBuffer.WriteString(",")
			// }
			// compliantBuffer.WriteString("{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
			// 	",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
			// 	",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\"" +
			// 	",\"source\":\"" + row.Columns[4].GetString_() + "\"}")
		} else if row.Columns[2].GetString_() == "non-compliant" {
			noncompliantCustomersCount++
			// 	if noncompliantBuffer.String() != "" {
			// 		noncompliantBuffer.WriteString(",")
			// 	}
			// 	noncompliantBuffer.WriteString("{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
			// 		",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
			// 		",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\"" +
			// 		",\"source\":\"" + row.Columns[4].GetString_() + "\"}")
		}
	}
	kycBuffer.WriteString("{" +
		"\"compliant\" : \"" + strconv.Itoa(compliantCustomersCount) + "\"," +
		"\"noncompliant\" : \"" + strconv.Itoa(noncompliantCustomersCount) + "\"," +
		"\"total\" : \"" + strconv.Itoa(totalCustomers) + "\"" +
		"}")

	bytes, err := json.Marshal(kycBuffer.String())
	if err != nil {
		return nil, errors.New("Error converting kyc stats")
	}
	return bytes, nil
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
