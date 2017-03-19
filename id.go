package main

import "math/rand"

const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

func id(n int) string {
	id := ""
	for i := 0; i < n; i++ {
		id += string(chars[rand.Intn(len(chars))])
	}
	return id
}
