package tests

import (
	"fmt"
	"github.com/bersen66/wb_task1/questions"
	"sync"
	"testing"
)

func TestPaginate(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	pages := questions.Paginate(slice, 1)
	fmt.Println(pages)
}

func TestComputing(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	var wg sync.WaitGroup
	for _, page := range questions.Paginate(slice, 2) {
		wg.Add(1)
		go questions.ComputeAndPrintSquares(slice, page, &wg)
	}
	wg.Wait()
}
