package questions

import (
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func DefineType(val interface{}) string {
	t := reflect.TypeOf(val)
	return t.String()
}
