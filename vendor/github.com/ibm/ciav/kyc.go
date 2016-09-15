/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package ciav

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Kyc struct {
	CustomerId  string
	KycStatus   string
	LastUpdated string
	Source      string
	KycRiskLevel string
}

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
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "riskLevel", Type: shim.ColumnDefinition_STRING, Key: false},
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
  var err error
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	var customerId string
	var kycStatus string
	var lastUpdated string
	var source string
	var riskLevel string

	isSuperAdmin, _ := stub.VerifyAttribute("role", []byte("Superadmin"))
	isManager, _ := stub.VerifyAttribute("role", []byte("Manager"))
	if isManager || isSuperAdmin {
		// customerId = args[0]
		// kycStatus = args[1]
		// lastUpdated = args[2]
		// source = args[3]
		// riskLevel = args[4]
		ok, err := stub.InsertRow("KYC", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
				&shim.Column{Value: &shim.Column_String_{String_: customerId}},
				&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: riskLevel}},
			},
		})

		if !ok && err == nil {
			return nil, errors.New("Error in adding KYC record.")
		}
	}else{
		// kycStr, _ := GetKYC(stub, args[0])
		// var kyc []Identification
		// err := json.Unmarshal([]byte(string(kycStr)), &kyc)
		// if err == nil {
		// 	// return nil, errors.New("Error in getting KYC record.")
		// }
		// customerId := args[0]
		// kycStatus := args[1]
		// lastUpdated := args[2]
		// source := args[3]
		// riskLevel = kyc.KycRiskLevel

		ok, err := stub.InsertRow("KYC", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
				&shim.Column{Value: &shim.Column_String_{String_: customerId}},
				&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: ""}},
			},
		})

		if !ok && err == nil {
			return nil, errors.New("Error in adding KYC record.")
		}
	}


	// myLogger.Debug("Congratulations !!! Successfully added")
	return nil, err
}

/*
	Update KYC record
*/
func UpdateKYC(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Updating KYC record ...")
	var err error
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}
		var customerId string
		var kycStatus string
		var lastUpdated string
		var source string
		var riskLevel string

		isSuperAdmin, _ := stub.VerifyAttribute("role", []byte("Superadmin"))
		isManager, _ := stub.VerifyAttribute("role", []byte("Manager"))
		if isManager || isSuperAdmin {
			ok, err := stub.InsertRow("KYC", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
					&shim.Column{Value: &shim.Column_String_{String_: customerId}},
					&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
					&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
					&shim.Column{Value: &shim.Column_String_{String_: source}},
					&shim.Column{Value: &shim.Column_String_{String_: riskLevel}},
				},
			})

			if !ok && err == nil {
				return nil, errors.New("Error in adding KYC record.")
			}
		}else{
			ok, err := stub.ReplaceRow("KYC", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
					&shim.Column{Value: &shim.Column_String_{String_: customerId}},
					&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
					&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
					&shim.Column{Value: &shim.Column_String_{String_: source}},
				},
			})

			if !ok && err == nil {
				return nil, errors.New("Error in adding KYC record.")
			}
		}
	// customerId := args[0]
	// kycStatus := args[1]
	// lastUpdated := args[2]
	// source := args[3]
	//
	// ok, err := stub.ReplaceRow("KYC", shim.Row{
	// 	Columns: []*shim.Column{
	// 		&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
	// 		&shim.Column{Value: &shim.Column_String_{String_: customerId}},
	// 		&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
	// 		&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
	// 		&shim.Column{Value: &shim.Column_String_{String_: "2"}},
	// 		&shim.Column{Value: &shim.Column_String_{String_: source}},
	// 	},
	// })
	//
	// if !ok && err == nil {
	// 	return nil, errors.New("Error in updating KYC record.")
	// }
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
	// callerRole, _ := stub.ReadCertAttribute("role")
	// jsonResp = jsonResp + ",\"role\":\"" + string(callerRole) + "\""
	jsonResp = jsonResp + ",\"source\":\"" + row.Columns[4].GetString_() + "\""

	isSuperAdmin, _ := stub.VerifyAttribute("role", []byte("Superadmin"))
	isManager, _ := stub.VerifyAttribute("role", []byte("Manager"))

	if isSuperAdmin || isManager {
		jsonResp = jsonResp + ",\"riskLevel\":\"" + row.Columns[5].GetString_() + "\""
	}else{
		jsonResp = jsonResp + ",\"riskLevel\":\"\""
	}
	jsonResp = jsonResp +"}"

	return jsonResp, nil
}
