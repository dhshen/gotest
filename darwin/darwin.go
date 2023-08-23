package darwin

import (
	"fmt"
	"github.com/dhshen/gotest"
)

func DarwinPrint(str string){
	fmt.Printf("darwin print %s\n", str)
	gotest.PrintSomething(str)
}
