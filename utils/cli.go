package utils

import "fmt"

func WelcomeUser() string {
	var name string
	for name == "" {
		fmt.Printf("May I have your name? ")
		fmt.Scanln(&name)
	}
	fmt.Printf("Hello, %s!\n", name)
	return name
}
