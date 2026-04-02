package games

import (
	"fmt"
	"strings"

	"github.com/kpechenenko/rword"
)

func Hangman() {
	var g rword.GenerateRandom
	var err error

	g, err = rword.New()
	if err != nil {
		panic(err)
	}

	lives := 5

	word := g.Word()
	revealed := make([]string, len(word))

	var userInput string

	flag := true
	guesses := make(map[string]bool)

	fmt.Printf("The word contains %d letters\n--------------\n", len(word))

	for flag {
		fmt.Println(revealed)
		fmt.Print("Guess a letter: ")
		fmt.Scan(&userInput)

		for i, char := range word {
			if string(char) == userInput {
				revealed[i] = userInput
			}
		}

		if guesses[strings.ToLower(userInput)] {
			fmt.Println("Already guessed that one!")
			continue
		}

		guesses[userInput] = true

		if strings.Contains(word, strings.ToLower(userInput)) {
			fmt.Printf("The word contains %s\n", userInput)
			fmt.Printf("You still have %d lives left\n-------------\n", lives)
		} else {
			fmt.Printf("The word doesn't contain %s\n", userInput)
			lives -= 1
			fmt.Printf("You lost a life! %d lives remaining\n-------------\n", lives)
		}
		if lives == 0 {
			flag = false
			fmt.Printf("You lose! The word was \"%s\"\n", word)
		}
	}
}
