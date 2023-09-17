package questions

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func ReverseWord(src string) string {
	repr := []rune(src)

	for i := 0; i < len(repr)/2; i++ {
		repr[i], repr[len(repr)-i-1] = repr[len(repr)-i-1], repr[i]
	}

	return string(repr)
}
