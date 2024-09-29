package infrastructure

import "fmt"

func findLetterIndices(word string, letter string) ([]int, error) {
	if len([]rune(letter)) != 1 {
		return nil, fmt.Errorf("В функцию переданы неверные значения")
	}

	target := []rune(letter)[0]
	wordChars := []rune(word)
	var indices []int
	for i, char := range wordChars {
		if char == target {
			indices = append(indices, i)
		}
	}
	return indices, nil
}

func insertSymbols(initialString string, letter string, indices []int) string {
	// Преобразуем строку в срез рун для возможности замены
	runes := []rune(initialString)

	// Преобразуем строку в руну (берём первый символ как руну, даже если это кириллица)
	letterRune := []rune(letter)[0]

	// Проходим по индексам
	for _, index := range indices {
		// Вычисляем позицию index*2
		pos := index * 2
		if pos < len(runes) { // Проверяем, что индекс не выходит за пределы строки
			runes[pos] = letterRune // Заменяем символ по индексу pos
		}
	}
	return string(runes)
}
