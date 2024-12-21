package main

import "database/sql"

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "chat.db")

	if err != nil {
		panic(err)
	}

	createTables()

	return db
}

func createTables() {
	db := GetConnection()

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		created DATETIME,
		is_deleted BOOLEAN
	);
	CREATE TABLE IF NOT EXISTS chats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		created DATETIME,
		closed DATETIME,
		is_closed BOOLEAN
	);
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chat_id INTEGER,
		user_id INTEGER,
		text TEXT,
		created DATETIME,
		is_deleted BOOLEAN
	);
	`

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}
