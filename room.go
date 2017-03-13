package main

import (
	"errors"

	"github.com/satori/go.uuid"
)

const maxRooms = 10

var rooms map[string]room
var roomService roomSvc

type roomSvc struct{}

type room struct {
	ID      string
	Content string
}

type roomAccess interface {
	New() (room, error)
	Update(id string, content string) bool
}

func (r roomSvc) New() (room, error) {
	if len(rooms) >= maxRooms {
		return room{}, errors.New("Too many open rooms")
	}
	newRoom := room{uuid.NewV4().String(), ""}
	if rooms == nil {
		rooms = make(map[string]room, maxRooms)
	}
	rooms[newRoom.ID] = newRoom
	return newRoom, nil
}

func (r roomSvc) Update(id string, content string) bool {
	_, ok := rooms[id]
	if ok {
		r := rooms[id]
		r.Content = content
		return true
	}
	return false
}
