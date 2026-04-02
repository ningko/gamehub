package games

import (
	"fmt"
	"math/rand"
)

func NumberGuesser() {
	num := 0 + rand.Intn(100-0)
	var userInput int
	guesses := make(map[int]bool)

	flag := true

	for flag {
		fmt.Print("Enter a random number: ")
		fmt.Scan(&userInput)
		if guesses[userInput] {
			fmt.Println("Already guessed that one!")
			continue
		}
		guesses[userInput] = true

		if userInput > num {
			fmt.Println("Lower!")
		} else if userInput < num {
			fmt.Println("Higher!")
		} else {
			flag = false
			fmt.Println("Correct!")
		}
	}

}
