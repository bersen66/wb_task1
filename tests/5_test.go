package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
	"time"
)

func Test5(t *testing.T) {
	questions.GenerateWithTimeout(1 * time.Second)
}
