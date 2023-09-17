package questions

import (
	"fmt"
	"math/rand"
)

// Реализовать паттерн «адаптер» на любом примере.

type Generator interface {
	Generate() int
}

type ConcreteGenerator struct {
	base int
}

func NewConcGen(base int) *ConcreteGenerator {
	return &ConcreteGenerator{
		base: base,
	}
}

func (g *ConcreteGenerator) GenerateInt() int {
	return rand.Int() % g.base
}

type ConcGenAdapter struct {
	g *ConcreteGenerator
}

func (c *ConcGenAdapter) Generate() int {
	return 0
}

func NewGenAdapter(g *ConcreteGenerator) *ConcGenAdapter {
	return &ConcGenAdapter{g: g}
}

func PrintGenerated(generator Generator) {
	fmt.Println(generator.Generate())
}
