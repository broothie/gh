package jog

import (
	"syscall/js"
)

type Generator interface {
	Generate() js.Value
}

type GeneratorFunc func() js.Value

func (f GeneratorFunc) Generate() js.Value {
	return f()
}
