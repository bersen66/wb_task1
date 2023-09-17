package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "абырвалг", questions.ReverseWord("главрыба"))
	assert.Equal(t, "йцукенгшщзхъ", questions.ReverseWord("ъхзщшгнекуцй"))
	assert.Equal(t, "量相步同", questions.ReverseWord("同步相量"))
	assert.Equal(t, "量相步同 ", questions.ReverseWord(" 同步相量"))
	assert.Equal(t, "qwe rty ", questions.ReverseWord(" ytr ewq"))
}
