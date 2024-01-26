package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getInput() string {
	// get input and use strings package to convert to lowercase b/c this is a good practice when working with inputs
	// in general
	fmt.Print("Pick [r]ock, [p]aper, or [s]cissors:   ")
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

var validInput = map[string]Move{
	"r":        Rock,
	"rock":     Rock,
	"p":        Paper,
	"paper":    Paper,
	"s":        Scissors,
	"scissors": Scissors,
}

var inputs = [...]string{
	Rock:     "Rock",
	Paper:    "Paper",
	Scissors: "Scissors",
}

func checkInput() Move {
	// start a loop until we get valid input
	for {
		// Does the user input match the map validInput?
		if move, ok := validInput[getInput()]; ok {
			fmt.Println("Player Chooses", inputs[move])
			return move
		}

		fmt.Println("I didn't understand your choice, please retry")
	}

}

var compare = [...]string{
	Rock:     "You Tied.",
	Paper:    "You Lost.",
	Scissors: "You Win!",
}

func main() {
	// for {} says do this until the user wants to stop with ctrl+c which is the universal way to kill a program
	// that's running in the terminal.
	// this is also known as 'while' loop in other languages and is how video games work (while user playing keep
	// running game engine)
	for {
		rand.Seed(time.Now().Unix())

		// Get Valid Player Input
		userChoice := checkInput()

		// Randomly Assign Computer Choice
		computerChoice := Move(rand.Intn(3))

		computerInput := inputs[computerChoice]

		fmt.Printf("Computer chooses: %s\n", computerInput)

		// Check to see who won
		fmt.Printf("Result: %s\n",
			compare[((computerChoice-userChoice)%3+3)%3])
	}

}
