package main

import "testing"

func TestSubscribeUnsubscribe(t *testing.T) {
	s1, err := ss.GetSubscribers("123")
	if err != nil {
		t.Error(err.Error())
	}
	if len(s1) != 0 {
		t.Errorf("unexpected number of subscriptions: %d", len(s1))
	}
	err = ss.Subscribe("123", subscription{"1", "me", nil})
	if err != nil {
		t.Error(err.Error())
	}
	s2, err := ss.GetSubscribers("123")
	if err != nil {
		t.Error(err.Error())
	}
	if len(s2) != 1 {
		t.Errorf("unexpected number of subscriptions: %d", len(s2))
	}
	err = ss.Unsubscribe("123", subscription{"1", "me", nil})
	if err != nil {
		t.Error(err.Error())
	}

	err = ss.Unsubscribe("456", subscription{"1", "me", nil})
	if err == nil {
		t.Error("should have failed due to non-existing room")
	}

	err = ss.Unsubscribe("123", subscription{"2", "someone-else", nil})
	if err == nil {
		t.Error("should have failed due to non-existing subscription")
	}
}

func TestMultipleSubscribers(t *testing.T) {
	roomID := "456"

	ss.Subscribe(roomID, subscription{"1", "me1", nil})
	if err := ss.Subscribe(roomID, subscription{"2", "me2", nil}); err != nil {
		t.Error(err.Error())
	}
	if err := ss.Subscribe(roomID, subscription{"3", "me3", nil}); err != nil {
		t.Error(err.Error())
	}

	subscribers, err := ss.GetSubscribers(roomID)
	if err != nil {
		t.Error(err.Error())
	}
	if len(subscribers) != 3 {
		t.Errorf("unexpected number of subscriptions: %d", len(subscribers))
	}
}
