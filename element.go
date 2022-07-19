package gh

import (
	"fmt"
	"strings"
	"syscall/js"
)

type Attr map[string]any

func Element(name string, attributes Attr, children ...Generator) GeneratorFunc {
	return func() js.Value {
		el := Document.CreateElement(name)

		for key, value := range attributes {
			if strings.HasPrefix(key, "on") {
				eventListener := value
				if listener, ok := value.(EventListener); ok {
					eventListener = js.FuncOf(func(_ js.Value, args []js.Value) any {
						return listener.HandleEvent(Event{Value: args[0]})
					})
				}

				el.Call("addEventListener", strings.TrimPrefix(key, "on"), eventListener)
			} else {
				el.Call("setAttribute", key, fmt.Sprint(value))
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
