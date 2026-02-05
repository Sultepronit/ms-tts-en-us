package db

import (
	"database/sql"
	"log"
	"time"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var conn *sql.DB

func Open() error {
	var err error
	conn, err = sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return err
	}

	conn.SetMaxOpenConns(1)
	conn.SetConnMaxLifetime(time.Hour)

	log.Println("DB opened")

	return nil
}
