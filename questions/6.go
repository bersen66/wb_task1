package questions

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func usingChan(ch chan int) {
	// Горутина блокируется в случае чтения из пустого канала
	val := <-ch
	fmt.Printf("After read: %v", val)
}

func usingWG() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		// Какие-то вычисления
	}(&wg)
	wg.Wait() // ожидание завершения задач
}

func usingMutex() {
	mutex := sync.Mutex{}
	mutex.Lock()
	// Если мьютекс захватить не удастся,то
	// горутина заблокируется до тех пор, пока
	// другая горутина не отпустит его
	defer mutex.Unlock()
}

func usingRWMutex() {
	mutex := sync.RWMutex{}
	mutex.Lock()
	// В случае, если кто-то читает данные, а мы
	// хотим их редактировать, то горутина будет заблокирована.
	// И наоборот.
	defer mutex.Unlock()
}

func usingTime() {
	time.Sleep(100 * time.Second) //просто спим 100 секунд
}

func usingScheduler() {
	runtime.Gosched() // Уходим в конец очереди планирования
}

func usingContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Другая логика
		}
	}
}

func usingExit() {
	os.Exit(0)
}

func usingPanic() {
	panic("Hello, world!")
}
