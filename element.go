package jog

import (
	"strings"
	"syscall/js"
)

type Attr map[string]any

func Element(name string, attributes Attr, children ...Generator) GeneratorFunc {
	return func() js.Value {
		el := document.Call("createElement", name)

		for key, value := range attributes {
			if strings.HasPrefix(key, "on") {
				if listener, ok := value.(EventListener); ok {
					el.Call("addEventListener", strings.TrimPrefix(key, "on"), js.FuncOf(func(_ js.Value, args []js.Value) any {
						return listener.HandleEvent(Event{JSValue: args[0]})
					}))
				} else {
					el.Call("addEventListener", strings.TrimPrefix(key, "on"), value)
				}
			} else {
				el.Call("setAttribute", key, value)
			}
		}

		for _, child := range children {
			el.Call("appendChild", child.Generate())
		}

		return el
	}
}

func Div(attributes Attr, children ...Generator) GeneratorFunc {
	return Element("div", attributes, children...)
}

func P(attributes Attr, children ...Generator) GeneratorFunc {
	return Element("p", attributes, children...)
}

func Input(attributes Attr, children ...Generator) GeneratorFunc {
	return Element("input", attributes, children...)
}
