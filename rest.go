package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/websocket"
)

func handleGetIndex(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		writeJSON(w, 400, err.Error())
		return
	}
	w.WriteHeader(200)
	io.WriteString(w, string(b))
}

func handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	room, err := roomService.New()
	if err != nil {
		writeJSON(w, 400, err.Error())
		return
	}
	w.WriteHeader(201)
	io.WriteString(w, room.ID)
}

func handleUpdateRoom(w http.ResponseWriter, r *http.Request) {
	var request roomUpdateRequest
	if err := parseJSON(r.Body, &request); err != nil {
		writeJSON(w, 400, err.Error())
		return
	}
	roomService.Update(request.ID, request.Update)

	subs, err := ss.GetSubscribers(request.ID)
	if err != nil {
		writeJSON(w, 400, err.Error())
		return
	}

	for _, sub := range subs {
		if request.UserID == sub.UserID {
			// skip this subscription - don't need to inform the user who posted this of the update.
			continue
		}
		if err = websocket.Message.Send(sub.Connection, request.Update); err != nil {
			// Don't fail the room update just because one send failed.
			fmt.Printf("Failed to deliver update to one of the subscribers: %s\n", err.Error())
			continue
		}
	}

	w.WriteHeader(204)
}

type roomUpdateRequest struct {
	ID     string `json:"roomId"`
	UserID string `json:"userId"`
	Update string `json:"update"`
}
