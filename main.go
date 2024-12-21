package main

import (
	models "alejandroblanco2001/chatroom/models"
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Isaac")
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chat")

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
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/chat", ChatHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
