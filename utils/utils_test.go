package utils

import (
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		num := GenerateRandomNumber()
		if num < 1 || num > 100 {
			t.Errorf("Generated number %d is out of range (1-100)", num)
		}
	}
}
