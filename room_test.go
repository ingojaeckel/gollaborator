package main

import "testing"

func TestRoomCreateUpdateGet(t *testing.T) {
	newRoom, err := roomService.New()
	if err != nil {
		t.Error(err.Error())
	}
	if newRoom.ID == "" {
		t.Error("new room id must not be empty")
	}
	if newRoom.Content != "" {
		t.Error("new room must be empty")
	}

	content, err := roomService.Get(newRoom.ID)
	if err != nil {
		t.Error(err.Error())
	}
	if content != "" {
		t.Error("new room must be empty")
	}

	if !roomService.Update(newRoom.ID, "hello world") {
		t.Error("room was not found")
	}

	content, err = roomService.Get(newRoom.ID)
	if err != nil {
		t.Error(err.Error())
	}
	if content != "hello world" {
		t.Errorf("room content update failed: %v", content)
	}
}

func TestRoomCreateTwoRooms(t *testing.T) {
	newRoom1, _ := roomService.New()
	newRoom2, _ := roomService.New()

	if newRoom1.ID == newRoom2.ID {
		t.Errorf("room ids must not match.")
	}
}
