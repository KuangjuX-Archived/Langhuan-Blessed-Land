package debug

import(
	"fmt"
)

func StdOutDebug(format string, params ...interface{}){
	format = "\033[1;34;40m" + format + "\033[0m\n"
	fmt.Printf(format, params...)
}