package repository

import (
	"database/sql"
	"log"
)

var db *sql.DB
func InitDatabaseRepository (receive_db *sql.DB) {
	db = receive_db
}

func VerifyApiToken () bool {
	row := db.QueryRow("SELECT id, api_key FROM user_data LIMIT 1")

	var id int
	var apiKey string

	err := row.Scan(&id, &apiKey)
	if err == sql.ErrNoRows {
		return false
  }
	if err != nil {
    log.Fatal(err)
  }

	return true
}

func InsertApiKey (apiKey string) {
	_, err := db.Exec("INSERT INTO user_data (api_key) VALUES (?)", apiKey)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetApiKey() (string, error) {
    var apiKey string

    err := db.QueryRow("SELECT api_key FROM user_data LIMIT 1").Scan(&apiKey)
    if err != nil {
        return "", err
    }

    return apiKey, nil
}
