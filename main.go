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

	games.IsEvenGame()
}
