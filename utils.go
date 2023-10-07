package main

import (
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
