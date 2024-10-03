package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

// LoadCategoriesFromJSON - Загрузка категорий из JSON-файла.
func LoadCategoriesFromJSON(filePath string) ([]domain.Category, error) {
	slog.Info("Loading categories from JSON file", slog.String("filePath", filePath))

	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Failed to open file", slog.String("filePath", filePath))
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			slog.Error("Failed to close file", slog.String("filePath", filePath))

			err = fmt.Errorf("ошибка при закрытии файла: %w", cerr)
		}
	}()

	// Читаем файл.
	data, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Failed to read file", slog.String("filePath", filePath))
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	var categoriesMap domain.CategoriesMap
	err = json.Unmarshal(data, &categoriesMap)

	if err != nil {
		slog.Error("Failed to parse JSON", slog.String("filePath", filePath))
		return nil, fmt.Errorf("ошибка при парсинге JSON: %w", err)
	}

	categories := make([]domain.Category, 0, len(categoriesMap))
	for categoryName, words := range categoriesMap {
		categories = append(categories, domain.Category{
			Name:  categoryName,
			Words: words,
		})
	}

	slog.Info("Categories successfully loaded", slog.Int("count", len(categories)))

	return categories, nil
}

// CategoryExists - Проверить, существует ли категория по названию.
func CategoryExists(categories []domain.Category, name string) bool {
	for _, category := range categories {
		if strings.EqualFold(category.Name, name) {
			slog.Info("Category found", slog.String("category", name))

			return true
		}
	}

	slog.Warn("Category not found", slog.String("category", name))

	return false
}
