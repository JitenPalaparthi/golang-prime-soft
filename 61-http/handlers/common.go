package handlers

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello PrimeSoft"))
	//w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	//w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
}
