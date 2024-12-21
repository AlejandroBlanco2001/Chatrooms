package models

import "time"

type Message struct {
	ID        int
	chatID    int
	userID    int
	text      string
	created   time.Time
	isDeleted bool
}
