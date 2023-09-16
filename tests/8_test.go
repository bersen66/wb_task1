package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetBit(t *testing.T) {
	var src int64 = 0
	questions.SetBit(&src, 0, true)
	assert.Equal(t, src, int64(1))
	questions.SetBit(&src, 1, true)
	assert.Equal(t, src, int64(3))
	questions.SetBit(&src, 2, true)
	assert.Equal(t, src, int64(3))
}

func TestInvalid(t *testing.T) {
	var src int64 = 0
	err := questions.SetBit(&src, 64, true)
	assert.NotEqual(t, err, nil)
	err = questions.SetBit(&src, -2, true)
	assert.NotEqual(t, err, nil)
}
