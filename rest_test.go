package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	len1 := len(rooms)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/room", nil)
	handleCreateRoom(rr, req)
	if rr.Code != 201 {
		t.Errorf("Unexpected status: %d", rr.Code)
	}

	len2 := len(rooms)

	if len2-len1 != 1 {
		t.Error("number of rooms has not increased by one")
	}

	id := rr.Body.String()
	_, ok := rooms[id]
	if !ok {
		t.Errorf("returned id of a room which does not exist: %s. rooms=%v", id, rooms)
	}
}

func TestUpdateRoom(t *testing.T) {
	rr1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/room", nil)
	handleCreateRoom(rr1, req1)
	roomID := rr1.Body.String()

	ss.Subscribe(roomID, subscription{UserID: "user-id", UserName: "me", Connection: nil})
	subs, _ := ss.GetSubscribers(roomID)
	if len(subs) != 1 {
		t.Errorf("unexpected number of subscriptions: %d", len(subs))
	}

	rr2 := httptest.NewRecorder()
	b, _ := toJSON(roomUpdateRequest{roomID, "user-id", "this is the new content"})
	req2, _ := http.NewRequest("UPDATE", "/room", bytes.NewReader(b))

	handleUpdateRoom(rr2, req2)
	if rr2.Code != 204 {
		t.Errorf("Unexpected status: %d", rr2.Code)
	}
}
