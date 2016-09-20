/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("customer_CIAV_details")

type ServicesChaincode struct {
}

func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("Hi Abhishek . . . ")
	err := stub.PutState("role", []byte("0"))
	return nil, err
}


func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	
	myLogger.Debug("I'm in Invoke . . . ")
	val, _ := stub.ReadCertAttribute("role")

	myLogger.Debug("Role : ", val)
	counter, _ := stub.PutState("role",val)
	
	return nil, nil

}

/*
 		Get Customer record by customer id or PAN number
*/
func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return read(stub, args)
}

func read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	
	role, _ := stub.GetState("role")
	
	jsonResp := "{ " + "Role: "+ role+"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	bytes, _ := json.Marshal(jsonResp)

	return bytes, nil
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
