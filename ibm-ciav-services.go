/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/chaincode"
	"github.com/op/go-logging"
	"github.com/hyperledger/fabric/core/chaincode/shim/crypto/attr"
)

var myLogger = logging.MustGetLogger("customer_CIAV_details")

type ServicesChaincode struct {
}

//type ChainName string
//const (
	// DefaultChain is the name of the default chain.
//	DefaultChain ChainName = "default"
//)

func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("Hi Abhishek . . . ")
	err := stub.PutState("role", []byte("0"))
	//var chains map[chaincode.ChainName]*chaincode.ChaincodeSupport
// 	chains = chaincode.GetChaincodeSupport()
// 	myLogger.Debug("Print chains . . . ")
// 	for key, val := range chains {
// 		myLogger.Debug("APP : Inside Loop . . . ")
// 	    	myLogger.Debug("key[%s] \n", key)
// 		myLogger.Debug("chains[name].peerAddress: %s", val)
// 	}
	//myLogger.Debug("Check Chain Name . . . ",chaincode.GetChainName())
	return nil, err
}


func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	myLogger.Debug("PEER ADDR: ",shim.GetPeerInfo())
	
	//var chain *chaincode.ChaincodeSupport
	myLogger.Debug("Initialize chains . . . ")
// 	var chains map[chaincode.ChainName]*chaincode.ChaincodeSupport

// 	myLogger.Debug("Assign value to chains . . . ")
// 	chains = chaincode.GetChaincodeSupport()
// 	myLogger.Debug("Print chains . . . ")
// 	for key, val := range chains {
// 		myLogger.Debug("APP INVOKE :Inside Loop . . . ")
// 	    	myLogger.Debug("key[%s] \n", key)
// 		myLogger.Debug("chains[name].peerAddress: %s", val)
// 	}
	
	
	myLogger.Debug("I'm in Invoke . . . ")
	//myLogger.Debug("Peer Id  ",stub.GetChain())
	// 6bf31ff0e07a759267344f84f97156b013189d1565c6d397c600decb64db5070b41e4e7dacf53e529e358bb56a83aaa206e8c1ee28b29b33bbb70777c5185a51
	//myLogger.Debug("Peer Id  ",chaincode.GetChain("default"))
	//myLogger.Debug("Peer Id  ",chain.peerAddress)
	
	val, _ := stub.ReadCertAttribute("role")
	sigma, _ := stub.GetCallerMetadata()
	payload, _ := stub.GetPayload()
	binding, _ := stub.GetBinding()
	
	adminCert, _ := stub.GetCallerMetadata()
	val1, _ := attr.GetValueFrom("role", adminCert)

	ok, _ := t.isCaller(stub, adminCert)
	if ok{
		myLogger.Debugf("passed verify . . .")
	}
	
//	myLogger.Debugf("passed certificate [% x]", certificate)
	myLogger.Debugf("passed sigma [% s]", string(sigma))
	myLogger.Debugf("passed payload [% s]", string(payload))
	myLogger.Debugf("passed binding [% s]", string(binding))
	
	myLogger.Debug("Role : ", string(val))
	myLogger.Debug("Role 1: ", string(val1))
	myLogger.Debug("~~~~~~~~~~~~~~~~~ END ~~~~~~~~~~~~~~~~")
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


