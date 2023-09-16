package questions

import "sync"

func Write(m *map[string]int, key string, val int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	(*m)[key] = val
}
