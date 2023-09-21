package main

import (
	"bufio"
	"os"
	"strings"
)

var wordList = make([]string, 6000)
var scanner = bufio.NewScanner(os.Stdin)

func loadWords() {
	wordsFile, err := os.ReadFile("words.txt")
	check(err)
	words := strings.Split(string(wordsFile), "\n")
	for i := 0; i < len(words); i++ {
		wordList[i] = string(words[i])
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
