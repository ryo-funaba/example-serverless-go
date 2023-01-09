package logutil

import (
	"fmt"
)

func PrintFuncName(structName, function string) {
	fmt.Printf("%s#%s\n", structName, function)
}
