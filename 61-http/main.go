package main

import (
	"http-demo/handlers"
	"log"
	"net/http"
	"runtime"
)

func main() {
	log.Println("The server is listening on port", port)
	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/user", handlers.CreatUser)

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Println(err.Error())
		runtime.Goexit() // this ensures that before going to crash, all ongoing goroutines gracefullu exit
		//log.Fatalln(err.Error())
	}
	//...
}
