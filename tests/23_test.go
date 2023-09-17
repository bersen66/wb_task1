package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErase(t *testing.T) {

	{
		src := []int{1, 2, 3, 4, 5}
		src = questions.Erase(src, 3)
		assert.Equal(t, []int{1, 2, 3, 5}, src)
	}

	{
		src := []int{1, 2, 3, 4, 5}
		src = questions.Erase(src, 0)
		assert.Equal(t, []int{2, 3, 4, 5}, src)
	}

	{
		src := []int{1, 2, 3, 4, 5}
		src = questions.Erase(src, len(src)-1)
		assert.Equal(t, []int{1, 2, 3, 4}, src)
	}
}
