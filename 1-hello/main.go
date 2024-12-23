package main // main is pre defined package

import (
	"fmt"
	"os"
	"time"
)

func main() { // main func...
	println("Hello Prime Soft")
	fmt.Println("Hello Prime Soft-->", time.Now())
	os.Exit(0)
}

func greet() {
	println("Hello Prime Soft")
}
