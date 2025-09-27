package main

import (
	"fmt"
	"main/utils"
)

func main() {
	fmt.Println(`
	Welcome to the Number Guessing Game!
	I'm thinking of a number between 1 and 100.
	You have 5 chances to guess the correct number.`)
	fmt.Println(`
	Please select the difficulty level:
	1. Easy (10 chances)
	2. Medium (5 chances)
	3. Hard (3 chances)`)
	fmt.Printf("Random number is %d\n", utils.GenerateRandomNumber())

	var user_choice int
	var difficulty = make(map[string]int, 0)
	difficulty["Easy"] = 10
	difficulty["Medium"] = 5
	difficulty["Hard"] = 3
	fmt.Println("Enter the first number.")
	fmt.Scanln(&user_choice)

	if user_choice <= 3 && user_choice >= 1 {
		fmt.Printf("User number is %d\n", user_choice)
	} else {
		fmt.Println("Invalid Choice")
	}

}
