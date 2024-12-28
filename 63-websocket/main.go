package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var mapConne map[string]*websocket.Conn

func init() {
	mapConne = make(map[string]*websocket.Conn)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for simplicity; restrict origins in production
		return true
	},
}

func main() {

	r := mux.NewRouter()
	// Start the server
	log.Printf("WebSocket server starting on 9096")
	http.Handle("/", r)
	r.HandleFunc("/ws", handleWebsocket)

	err := http.ListenAndServe(":9096", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer ws.Close()
	mapConne[fmt.Sprint(rand.IntN(1000))] = ws

	for {
		messageType, message, err := ws.ReadMessage()
		fmt.Println(messageType, string(message), err)

		for _, v := range mapConne {
			v.WriteMessage(1, message)
		}
	}

}
