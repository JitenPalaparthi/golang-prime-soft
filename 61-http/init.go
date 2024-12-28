package main

import (
	"flag"
	"os"
)

// var Global = func() int {
// 	return rand.IntN(100)
// }()

// func init() {
// 	log.Println(Global)
// }

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		flag.StringVar(&port, "port", "58089", "--port=58090 | -port=58089")
		flag.Parse()
	}
}
