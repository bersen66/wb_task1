package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"testing"
)

func Test1(t *testing.T) {

	h := questions.Human{
		Name:    "Ivan",
		Surname: "Ivanov",
		Age:     27,
	}

	ac := questions.Action{
		Human:      h,
		ActionName: "Jump",
	}

	if ac.IsHuman() != true {
		t.Fatal("ac is not human!")
	}

	ac.ActionMethod()
	ac.HumanMethod()

	if ac.Foo() != false {
		t.Fatal("Method not overrided")
	}

}
