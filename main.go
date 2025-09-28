package main

import (
	"fmt"
	"main/utils"
	"strconv"
	"time"
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
	random_number := utils.GenerateRandomNumber()
	fmt.Println(random_number)

	var user_choice int
	var difficulty = make(map[string]int, 0)
	difficulty["Easy"] = 10
	difficulty["Medium"] = 5
	difficulty["Hard"] = 3
	fmt.Println("Enter the first number.")
	fmt.Scanln(&user_choice)

	chances := 0
	if user_choice <= 3 && user_choice >= 1 {
		fmt.Printf("User number is %d\n", user_choice)
		switch user_choice {
		case 1:
			chances = difficulty["Easy"]
			fmt.Printf("Great! You have selected the Easy difficulty level.\nYou have %d chances\n", chances)
			fmt.Println("Let's start the game!")
		case 2:
			chances = difficulty["Medium"]
			fmt.Printf("Great! You have selected the Medium difficulty level.\nYou have %d chances\n", chances)
			fmt.Println("Let's start the game!")
		default:
			chances = difficulty["Hard"]
			fmt.Printf("Great! You have selected the Hard difficulty level.\nYou have %d chances\n", chances)
			fmt.Println("Let's start the game!")
		}
	} else {
		fmt.Println("Invalid Choice")
	}
	attempts := 0
	start_time := time.Now()

	for chances > 0 {
		var guess string
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if number, err := strconv.Atoi(guess); err == nil {
			if number < 101 && number > 0 {
				fmt.Printf("You have guessed %d\n", number)
				// break
			} else {
				fmt.Println("Number guess should be between 1 to 100")
				continue
			}
		} else {
			fmt.Println("Invalid input")
			continue
		}
		attempts++
		number, _ := strconv.Atoi(guess)
		if number == random_number {
			elapsed_time := time.Since(start_time)
			fmt.Printf("🎉 Congratulations! You guessed the number in %d attempts\n", attempts)
			fmt.Println("Time taken: ", elapsed_time)
			break
		} else if random_number < number {
			fmt.Printf("Incorrect! The number is less than %d\n", number)
		} else {
			fmt.Printf("Incorrect! The number is greater than %d\n", number)
		}
		chances--
		if chances == 0 {
			fmt.Printf("😞 You ran out of chances. The correct number was %d\n", random_number)
		}
	}
}
