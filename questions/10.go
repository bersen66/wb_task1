package questions

import (
	"math"
	"sort"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func groupOf(temp float64) int {
	val, _ := math.Modf(temp)
	res := int(val) / 10
	return res * 10
}

func Classify(in []float64) map[int][]float64 {
	result := make(map[int][]float64)

	sort.Float64s(in)
	for i := range in {
		group := groupOf(in[i])
		result[group] = append(result[group], in[i])
	}

	return result
}
