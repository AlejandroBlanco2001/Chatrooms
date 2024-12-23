package services

import (
	models "alejandroblanco2001/chatroom/models"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type ChatService struct {
	models.Chat
}

func (c *ChatService) FindOneChatByID(fields []string, chatId string) (models.Chat, error) {
	if len(fields) == 0 {
		fields = []string{"id", "name", "created", "closed", "is_closed"}
	}

	db := models.GetConnection()

	parsedFields := strings.Join(fields, ", ")

	query := `SELECT ` + parsedFields + ` FROM chats WHERE id = ?`

	stmt, err := db.Prepare(query)

	if err != nil {
		return models.Chat{}, err
	}

	defer stmt.Close()

	result, err := stmt.Query(chatId)

	if err != nil {
		return models.Chat{}, err
	}

	defer result.Close()

	var chat models.Chat

	if result.Next() {
		err = result.Scan(&chat.ID, &chat.Name, &chat.Created, &chat.Closed, &chat.IsClosed)

		if err != nil {
			return models.Chat{}, err
		}
	}

	return chat, nil
}
