package handlers

import (
	services "alejandroblanco2001/chatroom/services"
	"encoding/json"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Query().Get("id") != "" {
			SearchSpecificChatHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func SearchSpecificChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// IMPORTANT: I will not use the route /chat/:id because net/http does not support it
	// I will use the query parameter instead
	id := r.URL.Query().Get("id")

	chatService := services.ChatService{}
	chat, err := chatService.FindOneChatByID([]string{"id", "name", "created", "closed", "is_closed"}, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if chat.ID == 0 {
		http.Error(w, "Chat not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// send a JSOn with the chat

	if err := json.NewEncoder(w).Encode(chat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
