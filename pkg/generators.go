package pkg

import (
	"math/rand"
	"time"
)

func GenerateRandomLevel() int {
	// #nosec G404
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// #nosec G404
	return rand.Intn(4) // Генерирует число от 0 до 3.
}
