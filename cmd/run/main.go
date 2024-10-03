package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func main() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	absLogFilePath, err := filepath.Abs(viper.GetString("logPath"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile(absLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	infrastructure.WelcomeWord()
	session := infrastructure.StartSession()

	filePath := "internal/infrastructure/data/gamewords.json"
	categories, err := infrastructure.LoadCategoriesFromJSON(filePath)
	wordIsGiven := false

	if err != nil {
		slog.Error("Error loading categories", slog.String("filePath", filePath), slog.Any("error", err))
		fmt.Println("Ошибка при загрузке категорий:", err)

		return
	}

	session.Data = categories
	slog.Info("Categories loaded successfully", slog.Int("count", len(session.Data)))

	for {
		infrastructure.Writer(&session)
		infrastructure.Reader(&session)

		if session.SessionMode == 3 {
			break
		}

		if session.SessionMode == 1 && !wordIsGiven {
			word, found, gamefield := infrastructure.GenerateWord(session.Data, session.Category, session.Difficulty)

			if found {
				slog.Info("Random word generated", slog.String("word", word), slog.String("category", session.Category))
				fmt.Println("\n...Наколдовали вам случайное слово...")

				session.Word = word
				session.GameField = gamefield
				wordIsGiven = true
			} else {
				slog.Warn("No suitable word found", slog.String("category", session.Category), slog.Int("difficulty", session.Difficulty))
				fmt.Println("Подходящее слово не найдено. Попробуйте ввести другие требования.")

				session.SessionMode = 0
			}
		}

		session.SessionMode++
	}
}
