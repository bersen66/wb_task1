package questions

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// [Begin, End)
type Page struct {
	Begin int
	End   int
}

func (p *Page) Len() int {
	return p.End - p.Begin
}

func min(lhs, rhs int) int {
	if lhs <= rhs {
		return lhs
	}
	return rhs
}

func Paginate(slice []int, pageSize int) []Page {
	capacity := len(slice) / pageSize
	if len(slice)%pageSize != 0 {
		capacity++
	}
	result := make([]Page, 0, capacity)

	for i := 0; i < len(slice); i += pageSize {
		result = append(result, Page{
			Begin: i,
			End:   min(i+pageSize, len(slice)),
		})
	}

	return result
}

func ComputeAndPrintSquares(slice []int, page Page, group *sync.WaitGroup) {
	var val int
	for i := page.Begin; i < page.End; i++ {
		val = slice[i] * slice[i]
		fmt.Println(val)
	}
	group.Done()
}
