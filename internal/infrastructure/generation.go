package infrastructure

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func GenerateCategory(data []domain.Category) (string, bool) {
	if len(data) == 0 {
		slog.Warn("No categories available for random selection")
		return "", false
	}

	// Добавила данные комментарии, так как особая защита здесь не необходима.
	// #nosec G404
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// #nosec G404
	randomIndex := rand.Intn(len(data))

	slog.Info("Random category generated", slog.String("category", data[randomIndex].Name))

	return data[randomIndex].Name, true
}

func GenerateWord(data []domain.Category, categoryName string, difficulty int) (word string, fin bool, gameField string) {
	for _, category := range data {
		if !strings.EqualFold(category.Name, categoryName) {
			continue
		}

		var wordsByDifficulty []domain.Word

		for _, word := range category.Words {
			if word.Level == difficulty {
				wordsByDifficulty = append(wordsByDifficulty, word)
			}
		}

		if len(wordsByDifficulty) > 0 {
			// #nosec G404
			rand.New(rand.NewSource(time.Now().UnixNano()))
			// #nosec G404
			randomIndex := rand.Intn(len(wordsByDifficulty))
			gamefield := strings.TrimSpace(strings.Repeat("_ ", len([]rune(wordsByDifficulty[randomIndex].Word))))

			slog.Info("Word generated", slog.String("word", wordsByDifficulty[randomIndex].Word), slog.Int("difficulty", difficulty))

			return wordsByDifficulty[randomIndex].Word, true, gamefield
		}

		slog.Warn("No words found for specified difficulty", slog.String("category", categoryName), slog.Int("difficulty", difficulty))

		break
	}

	slog.Warn("Category not found", slog.String("categoryName", categoryName))

	return "", false, ""
}

func StartSession() domain.Session {
	session := domain.Session{
		SessionMode:    0,
		LastTriesCount: 0,
		Word:           "",
		Category:       "",
		LettersUsed:    make(map[rune]bool),
		Difficulty:     0,
		GameField:      "",
	}
	slog.Info("New game session started", slog.Any("session", session))

	return session
}
