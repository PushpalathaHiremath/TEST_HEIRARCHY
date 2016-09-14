package ROV_CIAV

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("Color")
type ServicesChaincode struct {
}

func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("Hello, My color is : , [%s]", GetColor())
	return nil, nil
}

func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return nil, nil
	
}
func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		myLogger.Debug("Hello, Error")
	//	fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}

