package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Brain Games!")

	var name string
	for name == "" {
		fmt.Printf("May I have your name? ")
		fmt.Scanln(&name)
	}
	fmt.Printf("Hello, %s!\n", name)
}
