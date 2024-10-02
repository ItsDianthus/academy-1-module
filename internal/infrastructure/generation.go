package infrastructure

import (
	"math/rand"
	"strings"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func GenerateCategory(data []domain.Category) (string, bool) {
	// Проверяем, есть ли категории в сессии.
	if len(data) == 0 {
		return "", false // Возвращаем пустую категорию, если нет категорий.
	}

	// Добавила данные комментарии, так как особая защита здесь не необходима.
	// #nosec G404
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// #nosec G404
	randomIndex := rand.Intn(len(data))

	return data[randomIndex].Name, true
}

func GenerateWord(data []domain.Category, categoryName string, difficulty int) (word string, fin bool, gameField string) {
	// Ищем категорию по имени
	for _, category := range data {
		if strings.ToLower(category.Name) == strings.ToLower(categoryName) {
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

				return wordsByDifficulty[randomIndex].Word, true, gamefield
			}

			break
		}
	}

	return "", false, ""
}

func StartSession() domain.Session {
	return domain.Session{
		SessionMode:    0,
		LastTriesCount: 0,
		Word:           "Жираф",
		Category:       "",
		LettersUsed:    make(map[rune]bool),
		Difficulty:     0,
		GameField:      "",
	}
}
