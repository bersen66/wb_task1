package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3(t *testing.T) {
	{
		src := []int{1, 2, 3, 4}
		result := questions.SumOfSquares(src, 4)
		assert.Equal(t, int64(30), result)
	}

	{
		src := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		result := questions.SumOfSquares(src, 4)
		assert.Equal(t, int64(0), result)
	}

}
