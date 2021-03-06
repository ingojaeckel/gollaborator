package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

func wsHandler(ws *websocket.Conn) {
	for {
		var err error
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		var subReq subscribeRequest
		if err = json.Unmarshal([]byte(reply), &subReq); err != nil {
			fmt.Println("Unable to parse request: ", err.Error())
			break
		}
		fmt.Printf("Received from client: %v\n", subReq)

		subscr := subscription{
			UserID:     subReq.UserID,
			UserName:   subReq.UserName,
			Connection: ws,
		}
		if err = ss.Subscribe(subReq.RoomID, subscr); err != nil {
			fmt.Println("Failed to subscribe: ", err.Error())
			break
		}
		fmt.Printf("Subscribed user %s to room %s\n", subReq.UserID, subReq.RoomID)

		// Send the current value as the first update to this client.
		initialValue, err := roomService.Get(subReq.RoomID)
		if err != nil {
			fmt.Println("Failed to retrieve room content: ", err.Error())
			break
		}
		if err = websocket.Message.Send(ws, initialValue); err != nil {
			fmt.Println("Failed to send out initial value: ", err.Error())
			break
		}
	}
	// Connection will be closed when reaching this point.
}

type subscribeRequest struct {
	RoomID   string `json:"roomId"`
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}
