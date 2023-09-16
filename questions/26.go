package questions

import "strings"

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

func IsUnique(src string) bool {
	freq := make(map[rune]int)

	for _, symb := range strings.ToLower(src) {
		if _, contains := freq[symb]; contains {
			return false
		}
		freq[symb]++
	}

	return true
}
