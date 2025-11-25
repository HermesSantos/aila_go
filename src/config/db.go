package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	execDir := filepath.Dir(execPath)
	fmt.Println("execDir", execDir)

	isGoRun := strings.Contains(execDir, "go-build")

	var dbPath string

	if isGoRun {
		wd, _ := os.Getwd()
		fmt.Println(wd)
		dbPath = filepath.Join(wd, "database.db")
	} else {
		dbPath = filepath.Join(execDir, "database.db")
	}

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user_data (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		api_key TEXT
	);`

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
