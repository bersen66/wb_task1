package questions

import "strings"

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
