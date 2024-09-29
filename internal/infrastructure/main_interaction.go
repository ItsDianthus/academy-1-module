package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Writer(session *domain.Session) {
	switch session.SessionMode {
	case domain.GetCategory:
		fmt.Print("1) Начнём с выбора категории. У нас есть: животные, профессии, наука.\nЕсли вы хотите случайную категорию, напишите '-'.\n")
	case domain.GetDifficulty:
		fmt.Print("2) Теперь необходимо выбрать сложность: число от 0 до 3, где 0 - новичок, а 3 - профессионал.\n")
	case domain.MainGame:
		fmt.Print("\nВаше слово: ")
		fmt.Println(session.GameField)
		fmt.Print("Использованные вами буквы: ")
		for letter := range session.LettersUsed {
			fmt.Printf("%c ", letter)
		}
		fmt.Print("\n")
		fmt.Println("Осталось шансов на ошибку: ", session.LastTriesCount)
		fmt.Print("Введите букву (русский алфавит): ")
	case domain.End:
		fmt.Print("Вы прошли игру.\n")
	default:
		fmt.Print("Всем привет. Это дефолт.\n")
	}
}

func Reader(session *domain.Session) {
	switch session.SessionMode {
	case domain.GetCategory:
		for {
			var category string
			_, err := fmt.Scan(&category)
			if err != nil {
				// Если произошла ошибка ввода, выводим сообщение
				fmt.Println("Попробуйте ввести категорию ещё раз")
			} else if category == "-" {
				category, found := GenerateCategory(*session)
				if found {
					fmt.Println("Случайная категория была выбрана:", category)
				} else {
					fmt.Println("Категории не найдены.")
				}
				break
			} else if !(CategoryExists(session.Data, category)) {
				// Если произошла ошибка ввода, выводим сообщение
				fmt.Println("Попробуйте ввести категорию ещё раз")
			} else {
				// Если ввод успешен, продолжаем
				fmt.Println("Вы ввели категорию: ", category)
				session.Category = category
				break
			}
		}

	case domain.GetDifficulty:
		scanner := bufio.NewScanner(os.Stdin)
		for {
			// Чтение ввода с проверкой на ошибку
			if scanner.Scan() {
				input := scanner.Text()

				if input == "-" {
					level := pkg.GenerateRandomLevel()
					session.Difficulty = level
					fmt.Println("Вам был сгенерирован уровень:", level)
					switch level {
					case 0:
						session.LastTriesCount = 12
					case 1:
						session.LastTriesCount = 12
					case 2:
						session.LastTriesCount = 6
					case 3:
						session.LastTriesCount = 4
					}
					break
				}
				// Пробуем преобразовать строку в число
				level, err := strconv.Atoi(strings.TrimSpace(input))
				if err != nil || level < 0 || level > 3 {
					fmt.Println("Ошибка: введите число от 0 до 3.")
				} else {
					session.Difficulty = level
					switch level {
					case 0:
						session.LastTriesCount = 12
					case 1:
						session.LastTriesCount = 12
					case 2:
						session.LastTriesCount = 6
					case 3:
						session.LastTriesCount = 4
					}
					break
				}
			}
		}
	case domain.MainGame:
		for {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]
			// Проверка на единственность буквы.
			if len([]rune(input)) != 1 {
				fmt.Println("Ошибка: введите ровно одну букву")

				// Проверка на кириллицу.
			} else if unicode.IsLetter([]rune(input)[0]) && unicode.In([]rune(input)[0], unicode.Cyrillic) {
				// Если это вообще русская буква
				if exists := session.LettersUsed[[]rune(input)[0]]; exists { // Если уже была использована
					fmt.Print("Буква ", input, " уже была использована. Введите другую букву: ")
				} else {
					indices, err := findLetterIndices(session.Word, input)
					if err != nil {
						fmt.Println("Ошибка:", err)

					} else if len(indices) != 0 {
						// fmt.Printf("Буква '%s' найдена на позициях: %s\n", input, strings.Trim(fmt.Sprint(indices), "[]"))
						// БУКВА УГАДАНА
						session.GameField = insertSymbols(session.GameField, input, indices)

						if !(strings.Contains(session.GameField, "_")) {
							EndGameWriter(true, session.Word)
							break
						}

					} else {
						// БУКВА НЕ УГАДАНА
						session.LastTriesCount -= 1
						HangmanWriter(session.LastTriesCount)
						if session.LastTriesCount == 0 {
							EndGameWriter(false, session.Word)
							break
						}
					}
					session.LettersUsed[[]rune(input)[0]] = true
					// Ещё не закончился mode MainGame
					session.SessionMode--
					break
				}
			} else {
				fmt.Println("Ошибка: введённый символ не является буквой русского алфавита")
			}
		}

	default:
		for {
			var category string
			_, err := fmt.Scan(&category)
			if err != nil {
				// Если произошла ошибка ввода, выводим сообщение
				fmt.Println("Ошибка при вводе: ", err)
				fmt.Println("Попробуйте ещё раз.")
			} else {
				// Если ввод успешен, продолжаем
				fmt.Println("Вы ввели категорию: ", category)
				break
			}
		}
	}
}
