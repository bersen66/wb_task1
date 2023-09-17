package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	{
		lhs := []int{1, 2, 3, 4, 5, 6}
		rhs := []int{1, 2, 3, 8, 9, 10}
		assert.Equal(t, questions.GetIntersection(lhs, rhs), []int{1, 2, 3})
	}
	{
		lhs := []int{}
		rhs := []int{1}
		assert.Equal(t, questions.GetIntersection(lhs, rhs), []int{})
	}

}
