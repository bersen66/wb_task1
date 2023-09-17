package questions

// Реализовать пересечение двух неупорядоченных множеств.

func GetIntersection(lhs, rhs []int) []int {
	result := make([]int, 0, len(lhs))

	searcher := make(map[int]bool)

	for i := range rhs {
		searcher[rhs[i]] = true
	}

	for i := range lhs {
		if _, cont := searcher[lhs[i]]; cont {
			result = append(result, lhs[i])
		}
	}

	return result
}
