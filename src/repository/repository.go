package repository

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

var db *sql.DB
func InitDatabase (receive_db *sql.DB) {
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

func SetCommitLanguage (commitLanguage string) (string, error) {
	_, err := db.Exec("UPDATE user_data SET commit_language = ?", commitLanguage)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "Idioma de commit atualizado", nil
}

func GetCommitLanguage() (string, error) {
	var commitLanguage string

	err := db.QueryRow("SELECT commit_language FROM user_data LIMIT 1").Scan(&commitLanguage)
	if err == nil {
		if commitLanguage == "" {
			commitLanguage = "english"
			_, _ = db.Exec(`UPDATE user_data SET commit_language = ?`, commitLanguage)
		}
		return commitLanguage, nil
	}

	if strings.Contains(err.Error(), "no such column") {
		_, alterErr := db.Exec(`ALTER TABLE user_data ADD COLUMN commit_language TEXT`)
		if alterErr != nil {
			return "", alterErr
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		_, insertErr := db.Exec(`INSERT INTO user_data (commit_language) VALUES ('english')`)
		if insertErr != nil {
			return "", insertErr
		}
		return "english", nil
	}

	return "", err
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
	if err == sql.ErrNoRows {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return apiKey, nil
}
