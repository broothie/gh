package jog

import (
	"context"
	"strings"
	"syscall/js"
)

type Attr map[string]any

type Element struct {
	Tag        string
	Attributes Attr
	Children   []Node

	subscriptions []string
}

func NewElement(tag string, attr Attr, children ...Node) Element {
	return Element{Tag: tag, Attributes: attr, Children: children}
}

func Div(attr Attr, children ...Node) Element {
	return NewElement("div", attr, children...)
}

func P(attr Attr, children ...Node) Element {
	return NewElement("p", attr, children...)
}

func Input(attr Attr, children ...Node) Element {
	return NewElement("input", attr, children...)
}

func (e Element) JSValue(ctx context.Context) js.Value {
	element := Document().Call("createElement", e.Tag)

	for key, value := range e.Attributes {
		if strings.HasPrefix(key, "on") {
			element.Call("addEventListener", strings.TrimPrefix(key, "on"), js.ValueOf(value))
		} else {
			element.Call("setAttribute", key, js.ValueOf(value))
		}
	}

	for _, child := range e.Children {
		element.Call("appendChild", child.JSValue(ctx))
	}

	jCtx := jogCtx(ctx)
	for _, subscription := range e.subscriptions {
		jCtx.subscriptions[subscription] = append(jCtx.subscriptions[subscription], element)
	}

	return element
}
