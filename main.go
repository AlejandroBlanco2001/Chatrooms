package main

import (
	handlers "alejandroblanco2001/chatroom/handlers"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Isaac")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/chat", handlers.ChatHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
