package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
)

func TestAdapter(t *testing.T) {
	conc := questions.NewConcGen(100)
	// questions.PrintGenerated(conc) Error - Incopatible interfaces!
	adapter := questions.NewGenAdapter(conc)
	questions.PrintGenerated(adapter) // OK, Compatible interfaces
}
