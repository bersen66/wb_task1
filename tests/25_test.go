package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	questions.Sleep(5 * time.Second)
}
