package main

import (
	"fmt"
	"math/rand"
)

var difficultyTry = map[string]int{
	"easy":   20,
	"medium": 15,
	"hard":   5,
}

var difficultyName = map[int]string{
	1: "easy",
	2: "medium",
	3: "hard",
}

func selectDifficulty() int {
	for {
		fmt.Printf("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

		fmt.Println("\n\nEnter your choice:")

		var difficulty int

		fmt.Scanln(&difficulty)

		if difficulty > 3 || difficulty <= 0 {
			fmt.Println("Invalid input, try again")
			continue
		}

		fmt.Printf("Great you have selectionned %s difficulty", difficultyName[difficulty])

		return difficultyTry[difficultyName[difficulty]]
	}
}

func main() {
	for {
		randomNumber := rand.Intn(100)

		fmt.Printf("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.")

		try := selectDifficulty()

		victory := false

		for try > 0 && !victory {
			fmt.Println("\n\nEnter your guess: ")

			var guess int

			fmt.Scanln(&guess)

			if guess == 0 || guess > 100 {
				fmt.Println("Invalid guess, check your input")
				continue
			}

			try--

			if guess == randomNumber {
				victory = true
				break
			}

			if guess < randomNumber {
				fmt.Printf("it's more than %d\n", guess)
			}

			if guess > randomNumber {
				fmt.Printf("it's less than %d\n", guess)
			}
		}

		if victory {
			fmt.Println("success !!!")
		} else {
			fmt.Printf("Too bad, you failed..., it was %d\n", randomNumber)
		}

		fmt.Println("\n\nDo you want to play again? [y/n]")

		var playAgain string
		fmt.Scanln(&playAgain)

		if playAgain == "y" || playAgain == "Y" {
			continue
		}

		fmt.Println("Thank you for playing! Bye...")
		return
	}
}
