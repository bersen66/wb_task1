package questions

import (
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42….) с использованием конкурентных вычислений.

func accumulate(seq []int, page Page, accum *int64, wg *sync.WaitGroup) {
	for i := page.Begin; i < page.End; i++ {
		val := int64(seq[i] * seq[i])
		atomic.AddInt64(accum, val)
	}
	wg.Done()
}

func SumOfSquares(seq []int, workers int) int64 {
	pages := Paginate(seq, workers)
	var accumulator int64 = 0

	var wg sync.WaitGroup
	for i := range pages {
		wg.Add(1)
		go accumulate(seq, pages[i], &accumulator, &wg)
	}
	wg.Wait()
	return accumulator
}
