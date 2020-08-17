package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"./api"
)

var addr = flag.String("addr", "localhost:4923", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func wsEntry(w http.ResponseWriter, r *http.Request) {
	wsClient, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	apiClient := api.CreateAPIClient(wsClient)
	defer apiClient.Close()

	apiClient.Hold()
}

func RunAPIServer() {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	http.HandleFunc("/v0/ws", wsEntry)

	log.Println("Server should be running at http://" + *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
