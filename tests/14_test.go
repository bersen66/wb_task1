package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test14(t *testing.T) {
	assert.Equal(t, "int", questions.DefineType(2))
	assert.Equal(t, "float64", questions.DefineType(3.2))
	assert.Equal(t, "bool", questions.DefineType(false))
	assert.Equal(t, "string", questions.DefineType("Hello, world"))
	assert.Equal(t, "chan int", questions.DefineType(make(chan int)))
	assert.Equal(t, "chan float64", questions.DefineType(make(chan float64)))
}
