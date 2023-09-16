package questions

import (
	"fmt"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func Printer(in chan int) {

LOOP:
	for {
		select {
		case val, ok := <-in:
			if !ok {
				break LOOP
			}
			fmt.Println(val)

		}
	}

}

func Modifier(in, out chan int) {

LOOP:
	for {
		select {
		case val, ok := <-in:
			if !ok {
				break LOOP
			}
			out <- (2 * val)
		}
	}

}

func RunConv(values []int) {
	toMod := make(chan int)
	toPrint := make(chan int)

	go Modifier(toMod, toPrint)
	go Printer(toPrint)

	for i := 0; i < len(values); i++ {
		toMod <- values[i]
	}
	close(toMod)

	time.Sleep(1 * time.Second)
}
