package questions

import (
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

func QuickSort(slice []int, compare func(lhs, rhs int) bool) {
	// Стандартная функция сортировки реализует алгоритм qsort
	sort.Slice(slice, compare)
	/*
		Реализация из стандартной библиотеки имеет вид:
			func Slice(x any, less func(i, j int) bool) {
				rv := reflectValueOf(x)
				swap := reflectSwapper(x)
				length := rv.Len()
				quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
			}
		В теле функции вызывается внутренняя, неимпортированная реализация quicksort
	*/
}
