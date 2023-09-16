package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
)

/*
	Тесты из условия
*/
func Test26(t *testing.T) {
	if questions.IsUnique("abcd") != true {
		t.Fatal("Unique str treats as not unique")
	}
	if questions.IsUnique("abCdefAaf") != false {
		t.Fatal("Not unique str treats as unique")
	}
	if questions.IsUnique("aabcd") != false {
		t.Fatal("Not unique str treats as unique")
	}
}

func Test26Empty(t *testing.T) {
	var src string = ""

	if questions.IsUnique(src) != true {
		t.Fatal("Unique str treats as not unique")
	}
}

func Test26NotUnique(t *testing.T) {
	var src string = "aaaaAAAA"

	if questions.IsUnique(src) != false {
		t.Fatal("Not unique str treats as unique")
	}
}
