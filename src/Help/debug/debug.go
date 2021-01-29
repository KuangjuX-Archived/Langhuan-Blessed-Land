package debug

import(
	"fmt"
)

func StdOut(format string, params ...interface{}){
	fmt.Printf(format, params...)
}