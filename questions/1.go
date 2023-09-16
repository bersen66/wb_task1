package questions

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h *Human) IsHuman() bool {
	return true
}

func (h *Human) HumanMethod() {
	return
}

func (h *Human) Foo() bool {
	return true
}

type Action struct {
	Human
	ActionName string
}

func (a *Action) ActionMethod() bool {
	return true
}

func (a *Action) Foo() bool {
	return false
}
