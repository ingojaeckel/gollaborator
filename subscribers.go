package main

import (
	"errors"

	"golang.org/x/net/websocket"
)

var ss subscriptionSvc

type subscriptionSvc struct {
	Subscribers map[string][]subscription
}

type subscription struct {
	UserID     string
	UserName   string
	Connection *websocket.Conn
}

type subscriber interface {
	All() map[string][]subscription
	Subscribe(roomID string, s subscription) error
	Unsubscribe(roomID string, s subscription) error
	GetSubscribers(roomID string) ([]subscription, error)
}

func (subSvc *subscriptionSvc) All() map[string][]subscription {
	return subSvc.Subscribers
}

func (subSvc *subscriptionSvc) Subscribe(roomID string, s subscription) error {
	if subSvc.Subscribers == nil {
		subSvc.Subscribers = make(map[string][]subscription)
	}
	existingSubs, ok := subSvc.Subscribers[roomID]
	if ok {
		// append to existing list of subscribers
		subSvc.Subscribers[roomID] = append(existingSubs, s)
	} else {
		// initialize list of subscribers
		subSvc.Subscribers[roomID] = []subscription{s}
	}
	return nil
}

func (subSvc *subscriptionSvc) GetSubscribers(roomID string) ([]subscription, error) {
	if subSvc.Subscribers == nil {
		return []subscription{}, nil
	}
	existingSubs, ok := subSvc.Subscribers[roomID]
	if ok {
		return existingSubs, nil
	}
	// there are no subscriptions for this room
	return []subscription{}, nil
}

func (subSvc *subscriptionSvc) Unsubscribe(roomID string, s subscription) error {
	if subSvc.Subscribers == nil {
		return nil
	}
	existingSubs, ok := subSvc.Subscribers[roomID]
	if !ok {
		return errors.New("Room not found")
	}
	for index, ex := range existingSubs {
		if s.Connection == ex.Connection && s.UserID == ex.UserID && s.UserName == ex.UserName {
			// found the subscription that should be removed. Copy original subscriptions but skip the current element.
			subSvc.Subscribers[roomID] = append(existingSubs[:index], existingSubs[index+1:]...)
			return nil
		}
	}
	return errors.New("Could not find the subscription")
}
