package ROV_CIAV

import (
	"fmt"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("customer_address_details")
func main() {
	fmt.Println("Hello, My color is : , [%s]", GetColor())
}

