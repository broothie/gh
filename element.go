package jog

import (
	"strings"
	"syscall/js"
)

func Element(name string, attributes Attr, children ...*Node) *Node {
	return &Node{
		generator: element{
			name:       name,
			attributes: attributes,
			children:   children,
		},
	}
}

func Div(attributes Attr, children ...*Node) *Node {
	return Element("div", attributes, children...)
}

func P(attributes Attr, children ...*Node) *Node {
	return Element("p", attributes, children...)
}

func Input(attributes Attr, children ...*Node) *Node {
	return Element("input", attributes, children...)
}

type Attr map[string]any

type element struct {
	name       string
	attributes Attr
	children   []*Node
}

func (e element) Generate() js.Value {
	el := document.Call("createElement", e.name)

	for key, value := range e.attributes {
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

	for _, child := range e.children {
		el.Call("appendChild", child.ToJSValue())
	}

	return el
}
