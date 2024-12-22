package models

import (
	"errors"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Chat representa una sala de chat en el sistema
type Chat struct {
	ID       int
	Name     string // Exportamos el campo para acceso externo
	Created  time.Time
	Closed   time.Time
	IsClosed bool
	Users    []User
}

// Create inserta un nuevo registro de chat en la base de datos
func (c *Chat) Create() error {
	db := GetConnection()

	query := `INSERT INTO chats (name, created, closed, is_closed) VALUES (?, ?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insertamos los valores
	result, err := stmt.Exec(c.Name, c.Created, c.Closed, c.IsClosed)
	if err != nil {
		return err
	}

	// Verificamos que la fila fue insertada
	if rows, err := result.RowsAffected(); err != nil || rows != 1 {
		return errors.New("error creating chat")
	}

	return nil
}

func FindOneChatByID(fields []string, chatId string) (Chat, error) {
	if len(fields) == 0 {
		fields = []string{"id", "name", "created", "closed", "is_closed"}
	}

	db := GetConnection()

	parsedFields := strings.Join(fields, ", ")

	query := `SELECT ` + parsedFields + ` FROM chats WHERE id = ?`

	stmt, err := db.Prepare(query)

	if err != nil {
		return Chat{}, err
	}

	defer stmt.Close()

	result, err := stmt.Query(chatId)

	if err != nil {
		return Chat{}, err
	}

	defer result.Close()

	var chat Chat

	if result.Next() {
		err = result.Scan(&chat.ID, &chat.Name, &chat.Created, &chat.Closed, &chat.IsClosed)

		if err != nil {
			return Chat{}, err
		}
	}

	return chat, nil
}
