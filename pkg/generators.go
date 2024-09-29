package pkg

import (
	"math/rand"
	"time"
)

func GenerateRandomLevel() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(4) // Генерирует число от 0 до 3
}
