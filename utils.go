package main

import (
	"bufio"
	"os"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

type Word struct {
	id   int
	name string
}
type Match struct {
	id          int
	wordToGuess string
	solved      bool
	tries       int
	result      string
	createdAt   time.Time
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
