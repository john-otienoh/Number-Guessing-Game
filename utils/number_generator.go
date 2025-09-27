package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := r.Intn(100) + 1
	return secretNumber
}
