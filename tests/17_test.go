package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBS(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	{
		i, ok := questions.BinarySearch(src, 1)
		assert.Equal(t, 0, i)
		assert.Equal(t, true, ok)
	}

	{
		i, ok := questions.BinarySearch(src, 0)
		assert.Equal(t, 0, i)
		assert.Equal(t, false, ok)
	}
}
