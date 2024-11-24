package games

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func IsEvenGame() bool {
	fmt.Println("Answer 'yes' if the number is even, otherwise answer 'no'.")

	wins := 0
	for i := 0; i < 3; i++ {
		randNum := rand.IntN(100)
		var correct string
		if randNum%2 == 0 {
			correct = "yes"
		} else {
			correct = "no"
		}
		fmt.Println("Question:", randNum)

		var answer string
		fmt.Scanf("%s", &answer)
		fmt.Println("Your answer is:", answer)

		answer = strings.TrimRight(answer, "\n")

		if answer == correct {
			wins += 1
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%s'.\n", answer, correct)
			return false
		}
	}
	return true
}
