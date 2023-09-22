package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	loadWords()

	for {
		fmt.Printf("Welcome to my Go Wordle! \nWhat you want to do? \n 1- Play a game \n 2- Read instructions \n 3- Exit \n")
		scanner.Scan()
		input := scanner.Text()
		inputInt, err := strconv.Atoi(input)
		check(err)
		switch inputInt {
		case 1:
			playGame()
		case 2:
			instructions()
		case 3:
			fmt.Println("Bye bye!")
		default:
			fmt.Println("Try again")
		}
		if inputInt == 3 {
			break
		}

	}

}

func checkCorrectInput(word string) error {
	if len(word) != 5 {
		return errors.New("word must have 5 characters")
	}
	for _, letter := range word {
		if rune(letter) <= 97 || rune(letter) >= 122 {
			return errors.New("word must only have letters")
		}
	}
	return nil
}

func compareWords(guess, word string) []int {
	result := make([]int, 5)

	for j := 0; j < 5; j++ {
		if guess[j] == word[j] {
			result[j] = 2
			word = strings.Replace(word, string(word[j]), "*", 1)
		}
	}
	for i := 0; i < 5; i++ {
		lettersInWord := strings.Count(word, string(guess[i]))
		lettersInGuess := strings.Count(string(guess[0:i]), string(guess[i]))
		if lettersInWord > 0 && lettersInGuess < lettersInWord {
			result[i] = 1
		} else if string(word[i]) != "*" {
			result[i] = 0
		}
	}

	return result
}

func playGame() {
	wordToGuess := wordList[rand.Intn(len(wordList))]
	for i := 5; i >= 0; i-- {
		fmt.Println("Guess the word:")
		scanner.Scan()
		guess := scanner.Text()
		err := checkCorrectInput(strings.ToLower(guess))
		if err == nil {
			compareResult := compareWords(guess, wordToGuess)

			yellow := color.New(color.FgYellow)
			green := color.New(color.FgGreen)
			correct := true

			for j := 0; j < 5; j++ {
				switch compareResult[j] {
				case 1:
					correct = false
					yellow.Print(string(guess[j]))
				case 2:
					green.Print(string(guess[j]))
				default:
					correct = false
					fmt.Print(string(guess[j]))
				}
				if j == 4 {
					fmt.Println()
				}
			}
			if correct {
				fmt.Println("Congratulations, you won!")
				break
			} else if i == 0 {
				fmt.Printf("You've lost, the word was %v\n", wordToGuess)
			} else {
				fmt.Printf("You have %v tries left\n", i)
			}
		} else {
			fmt.Println(err.Error())
			i++
		}

	}
}

func instructions() {
	fmt.Println("Wordle is a popular word puzzle game where the objective is to guess a hidden five-letter word within six attempts.")
}
