package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/exp/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

func PrintCategoryNames(session *domain.Session) {
	if len(session.Data) == 0 {
		slog.Warn("No categories available.")
		fmt.Println("Категории отсутствуют.")

		return
	}

	for _, category := range session.Data {
		fmt.Println("-", category.Name)
	}
}

func Writer(session *domain.Session) {
	switch session.SessionMode {
	case domain.GetCategory:
		fmt.Println("\n1) Начнём с выбора категории. Доступные категории: ")
		PrintCategoryNames(session)
		fmt.Println("Если вы хотите случайную категорию, напишите '-'")
	case domain.GetDifficulty:
		fmt.Println("\n2) Теперь необходимо выбрать сложность: число от 0 до 3, где 0 - новичок, а 3 - профессионал.")
	case domain.MainGame:
		fmt.Print("\nВаше слово: ")
		fmt.Println(session.GameField, "[ тема:", session.Category, "]")
		fmt.Print("Использованные вами буквы: ")

		for letter := range session.LettersUsed {
			fmt.Printf("%c ", letter)
		}

		fmt.Print("\n")
		fmt.Println("Осталось шансов на ошибку: ", session.LastTriesCount, "[ сложность:", session.Difficulty, "]")
		fmt.Print("Введите букву (русский алфавит): ")
	case domain.End:
		slog.Info("End of the game session.")
	default:
		fmt.Print("Всем привет. Это дефолт.\n")
	}
}

func HandleCategoryInput(session *domain.Session) {
	for {
		var category string
		_, err := fmt.Scan(&category)

		if err != nil {
			slog.Error("Error reading category input", slog.Any("error", err))
			fmt.Println("Попробуйте ввести категорию ещё раз")

			return
		}

		if category == "-" {
			category, found := GenerateCategory(session.Data)
			if found {
				slog.Info("Random category selected", slog.String("category", category))
				fmt.Println("Случайная категория была выбрана:", category)
				session.Category = category

				return
			}

			slog.Warn("No categories found for random selection.")
			fmt.Println("Категории не найдены.")

			return
		}

		if CategoryExists(session.Data, category) {
			slog.Info("User selected category", slog.String("category", category))
			fmt.Println("Вы ввели категорию: ", category)
			session.Category = category

			return
		}

		slog.Warn("Invalid category input", slog.String("input", category))
		fmt.Println("Попробуйте ввести категорию ещё раз")
	}
}

func HandleDifficultyInput(session *domain.Session) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			input := scanner.Text()

			if input == "-" {
				level := pkg.GenerateRandomLevel()
				SetDifficulty(level, session)
				slog.Info("Random difficulty level selected", slog.Int("difficulty", level))

				return
			}

			level, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil || level < 0 || level > 3 {
				slog.Warn("Invalid difficulty input", slog.String("input", input))
				fmt.Println("Ошибка: введите число от 0 до 3.")
			} else {
				SetDifficulty(level, session)
				slog.Info("User selected difficulty level", slog.Int("difficulty", level))

				return
			}
		}
	}
}

func Reader(session *domain.Session) {
	switch session.SessionMode {
	case domain.GetCategory:
		HandleCategoryInput(session)
	case domain.GetDifficulty:
		HandleDifficultyInput(session)
	case domain.MainGame:
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(input)
		input = input[:len(input)-1]

		symb, _ := utf8.DecodeRuneInString(input)

		defer func() {
			session.SessionMode--
		}()

		if len([]rune(input)) != 1 || (!unicode.IsLetter(symb) || !unicode.In(symb, unicode.Cyrillic)) {
			slog.Warn("Invalid character input", slog.String("input", input))
			fmt.Println("Ошибка: введённый символ не является буквой русского алфавита")

			return
		}

		if exists := session.LettersUsed[symb]; exists {
			slog.Warn("Letter already used", slog.String("letter", input))
			fmt.Println("Буква ", input, " уже была использована. Введите другую букву.")

			return
		}

		indices, err := findLetterIndices(session.Word, input)
		if err != nil {
			slog.Error("Error finding letter indices", slog.Any("error", err))
			fmt.Println("Ошибка:", err)

			return
		}

		if len(indices) != 0 {
			session.GameField = insertSymbols(session.GameField, input, indices)
			slog.Info("Letter guessed", slog.String("letter", input))

			if !strings.Contains(session.GameField, "_") {
				EndGameWriter(true, session.Word)
				slog.Info("User won the game", slog.String("word", session.Word))

				session.SessionMode++

				return
			}
		} else {
			session.LastTriesCount--
			HangmanWriter(session.LastTriesCount)
			slog.Warn("Letter not guessed", slog.String("letter", input))

			if session.LastTriesCount == 0 {
				EndGameWriter(false, session.Word)
				slog.Info("User lost the game", slog.String("word", session.Word))

				session.SessionMode++

				return
			}
		}

		session.LettersUsed[symb] = true
	case domain.End:
	}
}
