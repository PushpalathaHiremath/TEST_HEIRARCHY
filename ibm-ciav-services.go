/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
	"github.com/hyperledger/fabric/core/chaincode/shim/crypto/attr"
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
	sigma, _ := stub.GetCallerMetadata()
	payload, _ := stub.GetPayload()
	binding, _ := stub.GetBinding()
	
	adminCert, _ := stub.GetCallerMetadata()
	val1, _ := attr.GetValueFrom("role", adminCert)

//	myLogger.Debugf("passed certificate [% x]", certificate)
	myLogger.Debugf("passed sigma [% s]", string(sigma))
	myLogger.Debugf("passed payload [% s]", string(payload))
	myLogger.Debugf("passed binding [% s]", string(binding))
	
	myLogger.Debug("Role : ", string(val))
	myLogger.Debug("Role 1: ", string(val1))
	stub.PutState("role",val)

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

	jsonResp := "{ " + "Role: "+ string(role) +"}"
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
