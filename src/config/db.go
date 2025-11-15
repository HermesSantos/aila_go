package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error

	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user_data (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		api_key TEXT
	);
	`

	_, execErr := DB.Exec(sqlStmt)
	if execErr != nil {
		log.Fatal(execErr)
	}
}

func CloseDb () {
	if DB != nil {
		DB.Close()
	}
}
