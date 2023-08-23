package gotest

import (
	"fmt"
	"github.com/dhshen/gotest/darwin"
)

func PrintSomething(str string){
	fmt.Printf("v1.0.3 print str: %s.\n", str)
	darwin.DarwinPrint(str)
}