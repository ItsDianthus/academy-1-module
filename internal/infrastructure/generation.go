package infrastructure

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"math/rand"
	"strings"
	"time"
)

func GenerateCategory(session domain.Session) (string, bool) {
	// Проверяем, есть ли категории в сессии.
	if len(session.Data) == 0 {
		return "", false // Возвращаем пустую категорию, если нет категорий.
	}

	// Инициализируем генератор случайных чисел.
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Выбираем случайный индекс
	randomIndex := rand.Intn(len(session.Data))
	return session.Data[randomIndex].Name, true // Возвращаем случайную категорию.
}

func GenerateWord(session domain.Session, categoryName string, difficulty int) (string, bool, string) {
	// Ищем категорию по имени
	for _, category := range session.Data {
		if category.Name == categoryName {
			var wordsByDifficulty []domain.Word
			for _, word := range category.Words {
				if word.Level == difficulty {
					wordsByDifficulty = append(wordsByDifficulty, word)
				}
			}

			if len(wordsByDifficulty) > 0 {
				rand.New(rand.NewSource(time.Now().UnixNano()))
				// Выбираем случайный индекс
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
