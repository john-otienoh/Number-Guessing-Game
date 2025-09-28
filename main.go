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
	for {
		chances := utils.DifficultyLevel()
		utils.Game(chances)

		if !utils.PlayAgain() {
			fmt.Println("Thanks for playing! Goodbye")
			break
		}
	}
}
