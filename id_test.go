package main

import "testing"

func TestGetId(t *testing.T) {
	if len(id(5)) != 5 {
		t.Error("invalid id length")
	}
	if id(5) == id(5) {
		t.Error("two random ids must not match")
	}
}
