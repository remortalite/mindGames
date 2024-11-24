package main

import (
	"fmt"
	"mindGames/games"
	"mindGames/utils"
)

func main() {
	fmt.Println("Welcome to the Brain Games!")

	name := utils.WelcomeUser()
	fmt.Println(name)

	result := games.EvenGameRun()
	if result {
		fmt.Printf("Congratulations, %s!\n", name)
	} else {
		fmt.Printf("Let's try again, %s!\n", name)
	}
}
