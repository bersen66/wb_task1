package main

import (
	"fmt"
	"github.com/bersen66/wb_task1/questions"
)

func main() {

	fmt.Println("Enter number of workers:")
	var workers int
	if _, err := fmt.Scanf("%d", &workers); err != nil {
		fmt.Println("Invalid number")
		return
	} else {
		questions.StartWork(workers)
	}
}
