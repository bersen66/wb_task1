package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test7(t *testing.T) {
	m := make(map[string]int)
	expected := map[string]int{
		"Hello": 1,
		"World": 1,
		"Aboba": 13,
	}

	questions.Write(&m, "Hello", 1)
	questions.Write(&m, "Hello", 1)
	questions.Write(&m, "World", 1)
	questions.Write(&m, "Aboba", 13)

	assert.Equal(t, m, expected)
}
