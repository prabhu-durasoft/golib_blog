package utils

import (
	"math/rand"
	"time"
)

const PI = 3.1415

func GenerateRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
