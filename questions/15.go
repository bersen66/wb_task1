package questions

// 	К каким негативным последствиям может привести данный фрагмент кода, и как
//	это исправить? Приведите корректный пример реализации.
//
//	var justString string
//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//	}
//	func main() {
//		someFunc()
//	}

var justString string

func createHugeString(n int) string {
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = 'a'
	}
	return string(runes)
}

func someFunc() { // Bad!
	// В go строка представляет собой слайс байт,
	// однако, руны в UTF-8 могут занимать более одного байта.
	v := createHugeString(1 << 10)
	// В исходной реализации, срез осуществляется по байтам, а не по рунам.
	// Из-за этого срез может осуществиться некорректно
	justString = v[:100]
}

func coolSomeFunc() {
	v := createHugeString(1 << 10)
	// Конвертируем строку к слайсу рун
	// Срежем первые 100, как это сделано в образце
	// приведем срез к типу string,чтобы присвоить его justString
	justString = string([]rune(v)[:100])
}

//func main() {
//	someFunc()
//	fmt.Println(justString)
//}
