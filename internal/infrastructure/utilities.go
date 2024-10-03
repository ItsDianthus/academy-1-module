package infrastructure

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/exp/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func SetDifficulty(level int, session *domain.Session) {
	session.Difficulty = level

	switch level {
	case 0, 1:
		session.LastTriesCount = 12
	case 2:
		session.LastTriesCount = 6
	case 3:
		session.LastTriesCount = 4
	}

	slog.Info("Difficulty level set", slog.Int("level", level))
	fmt.Println("Установлен уровень сложности:", level)
}

func findLetterIndices(word, letter string) ([]int, error) {
	if len([]rune(letter)) != 1 {
		slog.Error("Invalid letter input", slog.String("letter", letter))
		return nil, fmt.Errorf("переданы неверные значения")
	}

	target, _ := utf8.DecodeRuneInString(letter)
	wordChars := []rune(word)

	var indices []int

	for i, char := range wordChars {
		if char == target {
			indices = append(indices, i)
		}
	}

	return indices, nil
}

func insertSymbols(initialString, letter string, indices []int) string {
	runes := []rune(initialString)

	letterRune, _ := utf8.DecodeRuneInString(letter)

	for _, index := range indices {
		pos := index * 2
		if pos < len(runes) {
			runes[pos] = letterRune
		}
	}

	return string(runes)
}
