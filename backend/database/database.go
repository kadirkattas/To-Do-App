package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

const databasePath = "./forum.db"

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}

	// Creating tables
	createTables := `
    CREATE TABLE IF NOT EXISTS USERS (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Email TEXT NOT NULL UNIQUE,
        UserName TEXT NOT NULL,
        Password TEXT NOT NULL,
		token TEXT
    );
    CREATE TABLE IF NOT EXISTS TODO (
		"ID" INTEGER UNIQUE,
		"UserID" INTEGER,
		"UserName" TEXT,
		"Title" TEXT,
		"Description" TEXT,
		"PostDate" TEXT,
		PRIMARY KEY("ID" AUTOINCREMENT)
    );`

	_, err = DB.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}
}
