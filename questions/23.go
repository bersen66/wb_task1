package questions

// Удалить i-ый элемент из слайса.

func Erase(slice []int, idx int) []int {
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	return slice[:len(slice)-1] // pop last
}
