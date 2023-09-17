package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

func Work(iterations int64, cnt *questions.Counter, wg *sync.WaitGroup) {
	for i := int64(0); i < iterations; i++ {
		cnt.Inc()
	}
	wg.Done()
}

func TestInc(t *testing.T) {
	cnt := questions.NewCounter()
	expected := int64(0)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		iterations := int64(rand.Int() % 1000)
		expected += iterations
		wg.Add(1)
		go Work(iterations, cnt, wg)
	}
	wg.Wait()
	assert.Equal(t, expected, cnt.Value())
}
