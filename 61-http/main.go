package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		flag.StringVar(&port, "port", "58089", "--port=58090 | -port=58089")
		flag.Parse()
	}

	fmt.Println("The server is listening on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello PrimeSoft"))
		//w.WriteHeader(200)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Println(err.Error())
		runtime.Goexit() // this ensures that before going to crash, all ongoing goroutines gracefullu exit
		//log.Fatalln(err.Error())
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

// what is the ,meaning of 0.0.0.0 --?

// /getMagicNumber
// it should give a random number which should be a primenumber
