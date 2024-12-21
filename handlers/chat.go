package handlers

import (
	"alejandroblanco2001/chatroom/models"
	"fmt"
	"net/http"
	"time"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat

	if r.Method == "POST" {
		chat.Name = "Chat de Isaac"
		chat.Created = time.Now()
		chat.IsClosed = false

		err := chat.Create()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	} else if r.Method == "GET" {
		row, err := chat.FindOneByID([]string{}, 3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Chat ID: %d, Name: %s, Created: %s, Closed: %s, IsClosed: %t",
			row.ID, row.Name, row.Created, row.Closed, row.IsClosed)
	}
}
