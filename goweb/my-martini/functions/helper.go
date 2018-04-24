package functions

import (
	"fmt"
	"os"
)

func Dump(name interface{}) {
	fmt.Println(name)
	os.Exit(0)
}
