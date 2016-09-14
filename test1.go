package ROV_CIAV

import (
	"fmt"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("Color")

func main() {
	myLogger.Debug("Hello, My color is : , [%s]", GetColor())
}

