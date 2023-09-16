package questions

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.

type WorkPool struct {
	size      int
	taskQueue chan int
	ctx       context.Context
}

func NewWorkPool(ctx context.Context, size int) *WorkPool {
	return &WorkPool{
		size:      size,
		taskQueue: make(chan int),
		ctx:       ctx,
	}
}

func Work(id int, taskChan <-chan int, ctx context.Context) {
LOOP:
	for {
		select {
		case data := <-taskChan:
			fmt.Printf("Worker %v: %v\n", id, data)
		case <-ctx.Done():
			break LOOP
		default:
			runtime.Gosched()
		}
	}
}

func (w *WorkPool) Run() {
	for i := 0; i < w.size; i++ {
		go Work( /*id=*/ i, w.taskQueue, w.ctx)
	}
}

func (w *WorkPool) Post(value int) {
	w.taskQueue <- value
}

func StartWork(workers int) {
	ctx, stopWorkers := context.WithCancel(context.Background())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	wp := NewWorkPool(ctx, workers)
	wp.Run()

LOOP:
	for {
		select {
		case acceptedSignal := <-signalChannel:
			if acceptedSignal == syscall.SIGINT {
				stopWorkers()
				break LOOP
			}
		default:
			wp.Post(rand.Int())
			time.Sleep(200 * time.Millisecond)
		}

	}
}
