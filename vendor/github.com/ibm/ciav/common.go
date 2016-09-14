/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package ciav

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var Superadmin map[string]string
var Manager map[string]string
var RelationalManager map[string]string
var Helpdesk map[string]string

func Initialize()() {
	Superadmin = map[string]string{
		"CustomerId":     "W",
		"IdentityNumber": "W",
		"PoiType":        "W",
		"PoiDoc":         "W",
		"ExpiryDate":     "W",
		"Source":         "W",
		"FirstName":      "W",
		"LastName":       "W",
		"Sex":            "W",
		"EmailId":        "W",
		"Dob":            "W",
		"PhoneNumber":    "W",
		"Occupation":     "W",
		"AnnualIncome":   "W",
		"IncomeSource":   "W",
		"KycStatus":      "W",
		"KycRiskLevel":   "W",
		"LastUpdated":    "W",
		"AddressId":      "W",
		"AddressType":    "W",
		"DoorNumber":     "W",
		"Street":         "W",
		"Locality":       "W",
		"City":           "W",
		"State":          "W",
		"Pincode":        "W",
		"PoaType":        "W",
		"PoaDoc":         "W"}

	Manager = map[string]string{
		"CustomerId":     "W",
		"IdentityNumber": "W",
		"PoiType":        "W",
		"PoiDoc":         "W",
		"ExpiryDate":     "W",
		"Source":         "W",
		"FirstName":      "W",
		"LastName":       "W",
		"Sex":            "W",
		"EmailId":        "W",
		"Dob":            "W",
		"PhoneNumber":    "W",
		"Occupation":     "W",
		"AnnualIncome":   "W",
		"IncomeSource":   "W",
		"KycStatus":      "W",
		"KycRiskLevel":   "W",
		"LastUpdated":    "W",
		"AddressId":      "W",
		"AddressType":    "W",
		"DoorNumber":     "W",
		"Street":         "W",
		"Locality":       "W",
		"City":           "W",
		"State":          "W",
		"Pincode":        "W",
		"PoaType":        "W",
		"PoaDoc":         "W"}

	RelationalManager = map[string]string{
		"CustomerId":     "W",
		"IdentityNumber": "W",
		"PoiType":        "W",
		"PoiDoc":         "W",
		"ExpiryDate":     "W",
		"Source":         "W",
		"FirstName":      "W",
		"LastName":       "W",
		"Sex":            "W",
		"EmailId":        "W",
		"Dob":            "W",
		"PhoneNumber":    "W",
		"Occupation":     "W",
		"AnnualIncome":   "W",
		"IncomeSource":   "W",
		"KycStatus":      "W",
		"KycRiskLevel":   "N",
		"LastUpdated":    "W",
		"AddressId":      "W",
		"AddressType":    "W",
		"DoorNumber":     "W",
		"Street":         "W",
		"Locality":       "W",
		"City":           "W",
		"State":          "W",
		"Pincode":        "W",
		"PoaType":        "W",
		"PoaDoc":         "W"}
	Helpdesk = map[string]string{
		"CustomerId":     "R",
		"IdentityNumber": "R",
		"PoiType":        "W",
		"PoiDoc":         "W",
		"ExpiryDate":     "R",
		"Source":         "W",
		"FirstName":      "R",
		"LastName":       "R",
		"Sex":            "R",
		"EmailId":        "R",
		"Dob":            "R",
		"PhoneNumber":    "R",
		"Occupation":     "R",
		"AnnualIncome":   "R",
		"IncomeSource":   "R",
		"KycStatus":      "R",
		"KycRiskLevel":   "N",
		"LastUpdated":    "R",
		"AddressId":      "R",
		"AddressType":    "R",
		"DoorNumber":     "R",
		"Street":         "R",
		"Locality":       "R",
		"City":           "R",
		"State":          "R",
		"Pincode":        "R",
		"PoaType":        "W",
		"PoaDoc":         "W"}
}

/*
	Get all rows corresponding to the partial keys given
*/
func GetAllRows(stub *shim.ChaincodeStub, tableName string, columns []shim.Column) ([]shim.Row, error) {
	rowChannel, err := stub.GetRows(tableName, columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving address details for : [%s]", err)
		return nil, fmt.Errorf("Failed retriving address details : [%s]", err)
	}
	var rows []shim.Row
	for {
		select {
		case temprow, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				// myLogger.Debugf("Fetching row : [%s]", temprow.Columns[0].GetString_())
				rows = append(rows, temprow)
			}
		}
		if rowChannel == nil {
			break
		}
	}
	return rows, nil
}

/*
 Get the customer id by PAN number
*/
func GetCustomerID(stub *shim.ChaincodeStub, panId string) ([]string, error) {
	var err error

	// myLogger.Debugf("Get customer id for PAN : [%s]", panId)

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: panId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("IDRelation", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Identification details for PAN [%s]: [%s]", string(panId), err)
		return nil, fmt.Errorf("Failed retriving Identification details  for PAN [%s]: [%s]", string(panId), err)
	}

	custIds := row.Columns[1].GetString_()
	custIdArray := strings.Split(custIds, "|")
	return custIdArray, nil
}
