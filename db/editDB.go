package db

import "log"

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Edit() {
	// createTable := `
	// DROP TABLE IF EXISTS voices;
	// CREATE TABLE IF NOT EXISTS voices (
	// id INTEGER PRIMARY KEY AUTOINCREMENT,
	// name TEXT NOT NULL,
	// code_name TEXT NOT NULL UNIQUE,
	// is_male BOOLEAN NOT NULL,
	// rate INTEGER NOT NULL DEFAULT -1,
	// rating INTEGER NOT NULL DEFAULT 0,
	// excluded BOOLEAN NOT NULL DEFAULT false,
	// comment TEXT NOT NULL DEFAULT ''
	// );
	// `
	createTable := `
	DROP TABLE IF EXISTS records;
	CREATE TABLE IF NOT EXISTS records (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	code_name TEXT NOT NULL UNIQUE,
	is_male BOOLEAN NOT NULL,
	rate INTEGER NOT NULL DEFAULT -1,
	rating INTEGER NOT NULL DEFAULT 0,
	excluded BOOLEAN NOT NULL DEFAULT false,
	comment TEXT NOT NULL DEFAULT ''
	);
	`
	_, err := conn.Exec(createTable)
	handleErr(err)

	log.Println("Successfully edited(?)")
}
