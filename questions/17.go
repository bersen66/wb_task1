package questions

import (
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func BinarySearch(slice []int, val int) (int, bool) {
	/*
		Стандартная функция поиска реализует алгоритм бинарного поиска.
			func Search(n int, f func(int) bool) int {
				// Define f(-1) == false and f(n) == true.
				// Invariant: f(i-1) == false, f(j) == true.
				i, j := 0, n
				for i < j {
					h := int(uint(i+j) >> 1) // avoid overflow when computing h
					// i ≤ h < j
					if !f(h) {
						i = h + 1 // preserves f(i-1) == false
					} else {
						j = h // preserves f(j) == true
					}
				}
				// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
				return i
			}
	*/
	i := sort.Search(len(slice), func(i int) bool { return slice[i] >= val })
	if i < len(slice) && slice[i] == val {
		return i, true
	} else {
		return i, false
	}
}
