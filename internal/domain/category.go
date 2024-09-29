package domain

type CategoriesMap map[string][]Word

type Category struct {
	Name  string `json:"-"`     // Название категории
	Words []Word `json:"words"` // Список слов
}
