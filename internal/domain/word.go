package domain

type Word struct {
	Word  string `json:"word"`  // Слово
	Level int    `json:"level"` // Уровень сложности
	Hint  string `json:"hint"`  // Подсказка
}
