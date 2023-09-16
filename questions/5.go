package questions

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"

	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func GenerateNumbers(out chan<- int, ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(200 * time.Millisecond)
			out <- rand.Int()
		}
	}
}

func GenerateWithTimeout(d time.Duration) {
	in := make(chan int)
	defer close(in)

	timer := time.NewTimer(d)

	ctx, finish := context.WithCancel(context.Background())

	go GenerateNumbers(in, ctx)

LOOP:
	for {
		select {
		case <-timer.C:
			finish()
			fmt.Println("Finished")
			break LOOP
		case val := <-in:
			fmt.Println("Generated:", val)
		default:
			runtime.Gosched()
		}
	}

}
