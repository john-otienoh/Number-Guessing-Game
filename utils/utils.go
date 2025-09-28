package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := r.Intn(100) + 1
	return secretNumber
}
func UserGuess() int {
	for {
		fmt.Print("Enter your guess (1-100): ")
		var input string
		fmt.Scanln(&input)

		if number, err := strconv.Atoi(input); err == nil {
			if number >= 1 && number <= 100 {
				return number
			}
			fmt.Println("Guess must be between 1 and 100.")
		} else {
			fmt.Println("Invalid input. Please enter a number.")
		}
	}
}
func PlayAgain() bool {
	var choice string
	fmt.Print("Do you want to play again? (y/n): ")
	fmt.Scanln(&choice)
	return choice == "y" || choice == "Y"
}

func DifficultyLevel() int {
	difficulty := map[int]struct {
		name    string
		chances int
	}{
		1: {"Easy", 10},
		2: {"Medium", 5},
		3: {"Hard", 3},
	}

	var choice int
	for {
		fmt.Print("Enter your choice: ")
		if _, err := fmt.Scanln(&choice); err == nil {
			if d, ok := difficulty[choice]; ok {
				fmt.Printf("Great! You have selected the %s difficulty level.\n", d.name)
				fmt.Printf("You have %d chances.\n", d.chances)
				fmt.Println("Let's start the game!")
				return d.chances
			}
		}
		fmt.Println("Invalid choice, please try again.")
	}
}

func Game(chances int) {
	randomNumber := GenerateRandomNumber()
	attempts := 0
	startTime := time.Now()

	for chances > 0 {
		guess := UserGuess()
		attempts++

		if guess == randomNumber {
			elapsed := time.Since(startTime)
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			fmt.Printf("Time taken: %s\n", elapsed.Round(time.Second))
			return
		}

		if guess > randomNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		}

		chances--
		if chances == 0 {
			fmt.Printf("You ran out of chances. The correct number was %d.\n", randomNumber)
		}
	}
}
