package questions

import "fmt"

// Поменять местами два числа без создания временной переменной.

func SwapAndPrint(lhs, rhs int) {
	fmt.Printf("Before: lhs=%v, rhs=%v\n", lhs, rhs)
	lhs += rhs
	rhs = lhs - rhs
	lhs = lhs - rhs
	fmt.Printf("After: lhs=%v, rhs=%v\n", lhs, rhs)
}
