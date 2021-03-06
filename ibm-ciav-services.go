/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
// 	"github.com/hyperledger/fabric/core/chaincode"
	"github.com/op/go-logging"
// 	"github.com/hyperledger/fabric/core/chaincode/shim/crypto/attr"
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
	myLogger.Debug("Hello ... Invoke")
// 	myLogger.Debug("Peer Id  ",stub.ChaincodeSupport.peerNetworkID)

	val, _ := stub.ReadCertAttribute("role")
	sigma, _ := stub.GetCallerMetadata()
	payload, _ := stub.GetPayload()
	binding, _ := stub.GetBinding()

	// adminCert, _ := stub.GetCallerMetadata()
	// val1, _ := attr.GetValueFrom("role", adminCert)

	ok, _ := t.isCaller(stub, sigma)
	if ok{
		myLogger.Debugf("passed verify . . .")
	}

//	myLogger.Debugf("passed certificate [% x]", certificate)
	myLogger.Debugf("Invoke passed sigma [% s]", string(sigma))
	myLogger.Debugf("Invoke passed payload [% s]", string(payload))
	myLogger.Debugf("Invoke passed binding [% s]", string(binding))

	myLogger.Debug("Invoke Role : ", string(val))
	myLogger.Debug("~~~~~~~~~~~~~~~~~Invoke END ~~~~~~~~~~~~~~~~")

	//var chain *chaincode.ChaincodeSupport
	// myLogger.Debug("Initialize chains . . . ")
	// 	var chains map[chaincode.ChainName]*chaincode.ChaincodeSupport
	// 	chains = chaincode.GetChaincodeSupport()
	// 	myLogger.Debug("Print chains . . . ")
	// 	for key, val := range chains {
	// 		myLogger.Debug("APP INVOKE :Inside Loop . . . ")
	// 	    	myLogger.Debug("key[%s] \n", key)
	// 		myLogger.Debug("chains[name].peerAddress: %s", val)
	// 	}
	//myLogger.Debug("Peer Id  ",stub.GetChain())
	//myLogger.Debug("Peer Id  ",chaincode.GetChain("default"))
	//myLogger.Debug("Peer Id  ",chain.peerAddress)
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

func (t *ServicesChaincode) isCaller(stub *shim.ChaincodeStub, certificate []byte) (bool, error) {
	myLogger.Debug("Check caller...")

	// In order to enforce access control, we require that the
	// metadata contains the signature under the signing key corresponding
	// to the verification key inside certificate of
	// the payload of the transaction (namely, function name and args) and
	// the transaction binding (to avoid copying attacks)

	// Verify \sigma=Sign(certificate.sk, tx.Payload||tx.Binding) against certificate.vk
	// \sigma is in the metadata

	sigma, err := stub.GetCallerMetadata()
	if err != nil {
		return false, err
	}
	payload, err := stub.GetPayload()
	if err != nil {
		return false, err
	}
	binding, err := stub.GetBinding()
	if err != nil {
		return false, err
	}

	myLogger.Debugf("passed certificate [% x]", certificate)
	myLogger.Debugf("passed sigma [% x]", sigma)
	myLogger.Debugf("passed payload [% x]", payload)
	myLogger.Debugf("passed binding [% x]", binding)

	ok, err := stub.VerifySignature(
		certificate,
		sigma,
		append(payload, binding...),
	)
	if err != nil {
		myLogger.Errorf("Failed checking signature [%s]", err)
		return ok, err
	}
	if !ok {
		myLogger.Error("Invalid signature")
	}

	myLogger.Debug("Check caller...Verified!")

	return ok, err
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
