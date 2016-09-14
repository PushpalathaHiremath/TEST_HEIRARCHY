/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package ciav

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				kyc
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*
	Create KYC table
*/
func CreateKycTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Creating KYC Table...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	err := stub.CreateTable("KYC", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "dummy", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "customerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "kycStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdated", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "riskLevel", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating KYC table.")
	}
	// myLogger.Debug("KYC table initialization done Successfully... !!! ")
	return nil, nil
}

/*
	Add KYC record
*/
func AddKYC(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Adding KYC record ...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	customerId := args[0]
	kycStatus := args[1]
	lastUpdated := args[2]
	source := args[3]

	ok, err := stub.InsertRow("KYC", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
			&shim.Column{Value: &shim.Column_String_{String_: "2"}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding KYC record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully added")
	return nil, err
}

/*
	Update KYC record
*/
func UpdateKYC(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Updating KYC record ...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	customerId := args[0]
	kycStatus := args[1]
	lastUpdated := args[2]
	source := args[3]

	ok, err := stub.ReplaceRow("KYC", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
			&shim.Column{Value: &shim.Column_String_{String_: "2"}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in updating KYC record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully updated")
	return nil, err
}

/*
 Get KYC record
*/
func GetKYC(stub *shim.ChaincodeStub, customerId string) (string, error) {
	var err error
	// myLogger.Debugf("Get identification record for customer : [%s]", string(customerId))
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: dummyValue}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: customerId}}
	columns = append(columns, col1)
	columns = append(columns, col2)
	row, err := stub.GetRow("KYC", columns)
	if err != nil {
		return "", fmt.Errorf("Failed retriving KYC details [%s]: [%s]", string(customerId), err)
	}
	jsonResp := "{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
		",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
		",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\""

	isOk, _ := stub.VerifyAttribute("role", []byte("Superadmin"))
	if isOk {
		jsonResp = jsonResp + ",\"riskLevel\":\"" + row.Columns[4].GetString_() + "\""
	}
	// callerRole, _ := stub.ReadCertAttribute("role")
	// jsonResp = jsonResp + ",\"role\":\"" + string(callerRole) + "\""
	jsonResp = jsonResp + ",\"source\":\"" + row.Columns[5].GetString_() + "\"}"

	return jsonResp, nil
}

func GetKYCStats(stub *shim.ChaincodeStub) ([]byte, error) {
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
