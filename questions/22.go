package questions

import "math/big"

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func AddBig(lhs, rhs *big.Int) *big.Int {
	var result big.Int
	result.Add(lhs, rhs)
	return &result
}

func SubBig(lhs, rhs *big.Int) *big.Int {
	var result big.Int
	result.Sub(lhs, rhs)
	return &result
}

func MulBig(lhs, rhs *big.Int) *big.Int {
	var result big.Int
	result.Mul(lhs, rhs)
	return &result
}

func DivBig(lhs, rhs *big.Int) *big.Int {
	var result big.Int
	result.Mul(lhs, rhs)
	return &result
}
