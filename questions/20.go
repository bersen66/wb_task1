package questions

import (
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func ReverseWords(text string) string {
	words := strings.Split(text, " ")

	builder := strings.Builder{}
	for i := range words {
		builder.WriteString(words[len(words)-i-1])
		if i != len(words)-1 {
			builder.WriteString(" ")
		}

	}

	//fmt.Println(words)
	return builder.String()
}
