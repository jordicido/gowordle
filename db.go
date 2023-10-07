package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func getConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "wordle.db")
	if err != nil {
		panic(err)
	}
	return db
}

func getWord() (string, error) {
	db := getConnection()

	var word struct {
		id   int
		name string
	}

	q := `SELECT * 
			FROM words 
			WHERE id = (ABS(random()) % (SELECT (SELECT MAX(id) FROM words)+1));`

	err := db.QueryRow(q).Scan(&word.id, &word.name)
	if err != nil {
		return "", err
	}

	return word.name, nil
}

func wordExists(word string) (bool, error) {
	db := getConnection()

	var Word struct {
		id   int
		name string
	}

	if err := db.QueryRow("SELECT * FROM words WHERE word = ?",
		word).Scan(&Word.id, &Word.name); err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("word " + word + " doesn't exists")
		}
		return false, errors.New("error retrieving data from database")
	}
	return true, nil
}
