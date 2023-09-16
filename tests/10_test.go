package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test10(t *testing.T) {

	expected := map[int][]float64{
		30:  {32.5},
		-20: {-27.0, -25.4, -21.0},
		10:  {13.0, 15.5, 19.0},
		20:  {24.5},
	}
	has := questions.Classify([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})
	assert.Equal(t, expected, has)
}
