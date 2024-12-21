package main

import "time"

type User struct {
	ID        int
	name      string
	email     string
	created   time.Time
	isDeleted bool
}
