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
	runes := []rune(initialString)

	letterRune := []rune(letter)[0]

	for _, index := range indices {
		pos := index * 2
		if pos < len(runes) {
			runes[pos] = letterRune
		}
	}
	return string(runes)
}
