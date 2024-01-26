// Game 1: Guess the Number
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Game: Guess a number between 0 and 10")
	fmt.Println("You have three(3) tries ")
	source := rand.NewSource(time.Now().UnixNano())

	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)
	var guess int

	for try := 1; try <= 3; try++ {
		fmt.Printf("Attempt: %d\n", try)
		fmt.Println("Enter your number: ")
		fmt.Scan(&guess)
		switch {
		case guess < secretNumber:
			fmt.Printf("Sorry, wrong GuessTheNumber ; number is too small\n ")
		case guess > secretNumber:
			fmt.Printf("Sorry, wrong GuessTheNumber ; number is too large\n ")
		case guess == secretNumber:
			fmt.Printf("YOU WIN!")
		case try == 3:
			fmt.Printf("Game over!\n ")
			fmt.Printf("The correct number is %d\n", secretNumber)
		default:
			fmt.Println("I don't recognize that input. Try a number like '1'")
		}

	}
}
