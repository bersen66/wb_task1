package tests

import (
	"github.com/bersen66/wb_task1/questions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {

	{
		set := questions.NewSet()
		src := []string{"cat", "cat", "dog", "cat", "tree"}
		exp := []string{"cat", "dog", "tree"}

		for i := range src {
			set.Insert(src[i])
		}

		i := 0
		for it := set.Iterator(); it.IsValid(); it.MoveNext() {
			assert.Equal(t, exp[i], it.Value())
			i++
		}
	}

	{
		set := questions.NewSet()
		src := []string{"a", "a", "a", "a"}
		exp := []string{"a"}

		for i := range src {
			set.Insert(src[i])
		}

		i := 0
		for it := set.Iterator(); it.IsValid(); it.MoveNext() {
			assert.Equal(t, exp[i], it.Value())
			i++
		}
	}

	{
		set := questions.NewSet()
		src := []string{"a", "b", "v", "c", "v", "c", "a"}
		exp := []string{"a", "b", "c", "v"}

		for i := range src {
			set.Insert(src[i])
		}
		i := 0
		for it := set.Iterator(); it.IsValid(); it.MoveNext() {
			assert.Equal(t, it.Value(), exp[i])
			i++

		}
	}

	{
		set := questions.NewSet()
		src := []string{}
		exp := []string{}

		for i := range src {
			set.Insert(src[i])
		}

		assert.Equal(t, set.Size(), len(exp))

	}

}
