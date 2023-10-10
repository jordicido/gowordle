package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {

	for {
		fmt.Printf("Welcome to my Go Wordle! \n\nWhat you want to do? \n\n 1- Play a game \n 2- Read instructions \n 3- Match history \n 4- Exit \n")
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
			showMatchHistory()
		case 4:
			fmt.Println("Bye bye!")
		default:
			fmt.Println("Try again")
		}
		if inputInt == 4 {
			break
		}

	}

}

func checkCorrectInput(word string) error {
	if len(word) != 5 {
		return errors.New("word must have 5 characters")
	}
	for _, letter := range word {
		if rune(letter) < 97 || rune(letter) > 122 {
			return errors.New("word must only have letters")
		}
	}
	_, err := wordExists(word)
	if err != nil {
		return err
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
	wordToGuess, err := getWord()
	check(err)
	var tries = [6]string{"", "", "", "", "", ""}
	correct := true
	for i := 0; i < 6; i++ {
		fmt.Println("Guess the word:")
		scanner.Scan()
		guess := scanner.Text()
		err := checkCorrectInput(strings.ToLower(guess))
		if err == nil {
			compareResult := compareWords(guess, wordToGuess)

			yellow := color.New(color.FgYellow)
			green := color.New(color.FgGreen)

			for j := 0; j < 5; j++ {
				switch compareResult[j] {
				case 1:
					correct = false
					tries[i] += yellow.Sprint(string(guess[j]))
				case 2:
					tries[i] += green.Sprint(string(guess[j]))
				default:
					correct = false
					tries[i] += fmt.Sprint(string(guess[j]))
				}
			}
			if correct {
				fmt.Println("Congratulations, you won!")
				fmt.Print("\n\n\n")
				break
			}
			fmt.Print("\n\n")
			for k := 0; k < 6; k++ {
				if len(tries[k]) > 0 {
					fmt.Println(tries[k])
				} else {
					break
				}
			}
			if i == 5 {
				fmt.Printf("You've lost, the word was %v\n", wordToGuess)
			} else {
				fmt.Printf("\nYou have %v tries left\n\n", 5-i)
			}
		} else {
			fmt.Println(err.Error())
			i--
		}

	}
	err = insertMatchResult(wordToGuess, tries[:], correct)
	check(err)
}

func instructions() {
	fmt.Println("Wordle is a popular word puzzle game where the objective is to guess a hidden five-letter word within six attempts.")
}

func showMatchHistory() {
	matchHistory, err := getMatchHistory()
	check(err)
	for _, match := range matchHistory {
		fmt.Printf("Partida: %d \n", match.id)
	}
}
