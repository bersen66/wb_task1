package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	src := []int{8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	questions.QuickSort(src, func(a, b int) bool { return src[a] < src[b] })

	assert.Equal(t, expected, src)
}
