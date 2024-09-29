package main

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"sync"
)

var once sync.Once

func main() {
	infrastructure.WelcomeWord()
	session := infrastructure.StartSession()
	wordIsGiven := false

	filePath := "internal/infrastructure/data/gamewords.json"
	categories, err := infrastructure.LoadCategoriesFromJSON(filePath)
	if err != nil {
		fmt.Println("Ошибка при загрузке категорий:", err)
		return
	}
	session.Data = categories

	//GenerateWord(session, )
	for {
		if session.SessionMode == 3 {
			fmt.Printf("Конец игры")
			break
		}
		infrastructure.Writer(&session)
		infrastructure.Reader(&session)
		if session.SessionMode == 1 && wordIsGiven == false {
			word, found, gamefield := infrastructure.GenerateWord(session, session.Category, session.Difficulty)
			if found {
				fmt.Println("Случайное слово было сгенерировано.")
				session.Word = word
				session.GameField = gamefield
				wordIsGiven = true
			} else {
				fmt.Println("Подходящее слово не найдено. Попробуйте ввести другие требования.")
				session.SessionMode = 0
			}
		}
		session.SessionMode++
	}
	return
}
