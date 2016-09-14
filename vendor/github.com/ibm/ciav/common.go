/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package ciav

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

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
