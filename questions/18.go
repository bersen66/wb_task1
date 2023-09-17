package questions

import "sync/atomic"

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое

type Counter struct {
	value int64
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
	}
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return c.value
}
