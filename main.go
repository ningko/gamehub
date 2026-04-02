package main

import (
	"fmt"
	"gamehub/games"
)

func main() {
	var choice string
	fmt.Println("Welcome to my shitty game hub:\n1) Number Guesser\n2) Hangman")
	fmt.Print("Select a game: ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		games.NumberGuesser()
	case "2":
		games.Hangman()
	default:
		fmt.Println("ok then bye")
	}
}
