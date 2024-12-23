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

func (c *Chat) FindByID(fields []string, chatId string) error {
	if len(fields) == 0 {
		fields = []string{"id", "name", "created", "closed", "is_closed"}
	}

	db := GetConnection()

	parsedFields := strings.Join(fields, ", ")

	query := `SELECT ` + parsedFields + ` FROM chats WHERE id = ?`

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Query(chatId)

	if err != nil {
		return err
	}

	defer result.Close()

	if result.Next() {
		err = result.Scan(&c.ID, &c.Name, &c.Created, &c.Closed, &c.IsClosed)

		if err != nil {
			return err
		}
	}

	return nil
}
