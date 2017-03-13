package main

import "testing"

func TestRoomCreate(t *testing.T) {
	newRoom, err := roomService.New()
	if err != nil {
		t.Error(err.Error())
	}
	if newRoom.ID == "" {
		t.Error("new room id must not be empty")
	}
	_, ok := rooms[newRoom.ID]
	if !ok {
		t.Error("room did not actually get created")
	}
}

func TestRoomCreateTwoRooms(t *testing.T) {
	newRoom1, _ := roomService.New()
	newRoom2, _ := roomService.New()

	if newRoom1.ID == newRoom2.ID {
		t.Errorf("room ids must not match.")
	}
}
