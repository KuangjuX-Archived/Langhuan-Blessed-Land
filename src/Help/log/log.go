package log

import (
	"fmt"
	"os"
	"time"
	"runtime"

	"github.com/sirupsen/logrus"
)


func Log(msg string){
	current_time := time.Now().Format("2006-01-02")
	var filename string
	sys := runtime.GOOS 
	switch sys {
		case "windows":
			filename = current_time + ".txt"
		case "macos":
		case "linux":
			filename = current_time	
	}
	path := "storage/logs/" + filename
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil{
		fmt.Printf("Open file %v failed", path)
		return
	}
	logrus.SetOutput(file)
	logrus.Info(msg)

}

