package util

import (
	"math/rand"
	"time"
)

// GenerateRandomNumber generates a random number for unique slug generation
func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999) + 1000
}
