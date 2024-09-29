package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"io"
	"os"
)

// LoadCategoriesFromJSON - Загрузка категорий из JSON-файла.
func LoadCategoriesFromJSON(filePath string) ([]domain.Category, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = fmt.Errorf("ошибка при закрытии файла: %w", cerr)
		}
	}()

	// Читаем файл.
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	// Парсим JSON в структуру.
	var categoriesMap domain.CategoriesMap
	err = json.Unmarshal(data, &categoriesMap)
	if err != nil {
		return nil, fmt.Errorf("ошибка при парсинге JSON: %w", err)
	}

	// Преобразуем CategoriesMap в массив Category.
	var categories []domain.Category
	for categoryName, words := range categoriesMap {
		categories = append(categories, domain.Category{
			Name:  categoryName,
			Words: words,
		})
	}

	return categories, nil
}

// CategoryExists - Проверить, существует ли категория по названию.
func CategoryExists(categories []domain.Category, name string) bool {
	for _, category := range categories {
		if category.Name == name {
			return true
		}
	}
	return false
}
