package main

import (
	"database/sql"
	"errors"
	"strings"

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

	var word Word

	q := `SELECT * 
			FROM words 
			WHERE id = (ABS(random()) % (SELECT (SELECT MAX(id) FROM words)+1));`

	err := db.QueryRow(q).Scan(&word.id, &word.name)
	if err != nil {
		return "", err
	}

	return word.name, nil
}

func wordExists(name string) (bool, error) {
	db := getConnection()

	var word Word

	if err := db.QueryRow("SELECT * FROM words WHERE word = ?",
		name).Scan(&word.id, &word.name); err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("word " + name + " doesn't exists")
		}
		return false, errors.New("error retrieving data from database")
	}
	return true, nil
}

func insertMatchResult(wordToGuess string, wordsTried []string, solved bool) error {
	db := getConnection()

	stmt, err := db.Prepare("INSERT INTO match_history (word, solved, tries, result) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(wordToGuess, solved, len(wordsTried), strings.Join(wordsTried, ","))
	if err != nil {
		return err
	}

	return nil
}

func getMatchHistory() ([]Match, error) {
	db := getConnection()
	var result []Match
	rows, err := db.Query("SELECT * FROM match_history")
	if err != nil {
		return nil, errors.New("error retrieving data from database")
	}
	defer rows.Close()

	for rows.Next() {
		var match Match
		err := rows.Scan(&match.id, &match.wordToGuess, &match.solved, &match.tries, &match.result, &match.createdAt)
		if err != nil {
			return nil, err
		}
		result = append(result, match)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
