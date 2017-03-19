package main

import "errors"

var roomService roomSvc

type roomSvc struct {
	Rooms map[string]*room
}

type room struct {
	ID      string
	Content string
}

type roomAccess interface {
	New() (room, error)
	Get(id string) (string, error)
	Update(id string, content string) bool
}

func (rs *roomSvc) New() (room, error) {
	if len(rs.Rooms) >= maxRooms {
		return room{}, errors.New("Too many open rooms")
	}
	newRoom := room{id(5), ""}
	if rs.Rooms == nil {
		rs.Rooms = make(map[string]*room, maxRooms)
	}
	rs.Rooms[newRoom.ID] = &newRoom
	return newRoom, nil
}

func (rs *roomSvc) Update(id string, content string) bool {
	r, ok := rs.Rooms[id]
	if !ok {
		return false
	}
	r.Content = content
	return true
}

func (rs roomSvc) Get(id string) (string, error) {
	if rs.Rooms == nil {
		return "", errors.New("Room not found")
	}

	r, ok := rs.Rooms[id]
	if !ok {
		return "", errors.New("Room not found")
	}
	return r.Content, nil
}
