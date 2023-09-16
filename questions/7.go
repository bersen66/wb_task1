package questions

import "sync"

// Реализовать конкурентную запись данных в map.

func Write(m *map[string]int, key string, val int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	(*m)[key] = val
}
