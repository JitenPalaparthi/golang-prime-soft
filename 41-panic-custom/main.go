package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// if strings.Contains(err.Error(), "no such file or directory") {
		// 	f, err := os.Create("data.txt")
		// }
		println(err.Error())
		panic("file is not available-->" + err.Error())
	}
	f.Close()
}

// any non zero exit is considered as a failure exit of the application.
//
