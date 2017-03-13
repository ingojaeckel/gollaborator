package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const portHTTP = 8080
const portWebsocket = 8081

func main() {
	mxWS := http.NewServeMux()
	mxWS.Handle("/", websocket.Handler(wsHandler))
	go func() {
		fmt.Println("Opening Websocket port ", portWebsocket)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portWebsocket), mxWS))
	}()

	mxHTTP := http.NewServeMux()
	mxHTTP.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetIndex(w, r)
		} else {
			w.WriteHeader(415)
		}
	})
	mxHTTP.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleCreateRoom(w, r)
		} else if r.Method == "UPDATE" {
			handleUpdateRoom(w, r)
		} else {
			w.WriteHeader(415)
		}
	})
	fmt.Println("Opening HTTP port ", portHTTP)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portHTTP), mxHTTP))
}
