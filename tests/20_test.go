package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "world Hello", questions.ReverseWords("Hello world"))
	assert.Equal(t, "frog", questions.ReverseWords("frog"))
	assert.Equal(t, "frog", questions.ReverseWords("frog"))
	assert.Equal(t, "sun dog snow", questions.ReverseWords("snow dog sun"))
}
