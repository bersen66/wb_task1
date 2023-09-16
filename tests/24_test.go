package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
)

func TestDistances(t *testing.T) {
	a := questions.NewPoint(0, 0)
	b := questions.NewPoint(3, 4)

	if a.DistanceTo(b) != 5 {
		t.Fatalf("Distance computed incorrectly, has %v, expected %v\n", a.DistanceTo(b), 5)
	}

	if b.DistanceTo(a) != 5 {
		t.Fatalf("Distance computed incorrectly, has %v, expected %v\n", b.DistanceTo(a), 5)
	}
}
