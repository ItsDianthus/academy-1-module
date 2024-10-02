package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

func PrintCategoryNames(session *domain.Session) {
	if len(session.Data) == 0 {
		fmt.Println("Категории отсутствуют.")

		return
	}

	fmt.Println("Доступные категории:")

	for _, category := range session.Data {
		fmt.Println("-", category.Name)
	}
}

func Writer(session *domain.Session) {
	switch session.SessionMode {
	case domain.GetCategory:
		fmt.Println("\n1) Начнём с выбора категории.")
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
		// Здесь мог бы быть вопрос про выход из игры.
	default:
		fmt.Print("Всем привет. Это дефолт.\n")
	}
}

func HandleCategoryInput(session *domain.Session) {
	for {
		var category string
		_, err := fmt.Scan(&category)

		if err != nil {
			// Если произошла ошибка ввода, выводим сообщение
			fmt.Println("Попробуйте ввести категорию ещё раз")
			return
		}

		if category == "-" {
			// Генерация случайной категории
			category, found := GenerateCategory(session.Data)
			if found {
				fmt.Println("Случайная категория была выбрана:", category)
				session.Category = category

				return
			}

			fmt.Println("Категории не найдены.")

			return
		}

		if CategoryExists(session.Data, category) {
			// Если ввод успешен, продолжаем
			fmt.Println("Вы ввели категорию: ", category)
			session.Category = category

			return
		}

		fmt.Println("Попробуйте ввести категорию ещё раз")
	}
}

func HandleDifficultyInput(session *domain.Session) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			input := scanner.Text()

			if input == "-" {
				// Генерация случайного уровня сложности
				level := pkg.GenerateRandomLevel()
				SetDifficulty(level, session)

				return
			}

			// Пробуем преобразовать строку в число
			level, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil || level < 0 || level > 3 {
				fmt.Println("Ошибка: введите число от 0 до 3.")
			} else {
				SetDifficulty(level, session)

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

		// Проверка на единственность буквы.
		if len([]rune(input)) != 1 {
			fmt.Println("Ошибка: введите ровно одну букву")
			return
		}

		if !unicode.IsLetter(symb) || !unicode.In(symb, unicode.Cyrillic) {
			fmt.Println("Ошибка: введённый символ не является буквой русского алфавита")
			return
		}

		if exists := session.LettersUsed[symb]; exists {
			fmt.Println("Буква ", input, " уже была использована. Введите другую букву.")
			return
		}

		indices, err := findLetterIndices(session.Word, input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		if len(indices) != 0 {
			// БУКВА УГАДАНА
			session.GameField = insertSymbols(session.GameField, input, indices)

			if !strings.Contains(session.GameField, "_") {
				EndGameWriter(true, session.Word)

				session.SessionMode++

				return
			}
		} else {
			// БУКВА НЕ УГАДАНА
			session.LastTriesCount--
			HangmanWriter(session.LastTriesCount)

			if session.LastTriesCount == 0 {
				EndGameWriter(false, session.Word)

				session.SessionMode++

				return
			}
		}

		session.LettersUsed[symb] = true
	case domain.End:
	// Здесь мог бы быть функционал считывания кнопки Enter/Escape для продолжения/остановки игры.
	default:
		fmt.Println("Ошибка.")
	}
}
