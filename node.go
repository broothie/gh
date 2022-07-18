package jog

import "syscall/js"

type Node struct {
	current   *js.Value
	generator Generator
}

func (n *Node) ToJSValue() js.Value {
	value := n.generator.Generate()
	n.current = &value
	return value
}

func (n *Node) update(newNode *Node) {
	oldValue := n.current
	newValue := newNode.ToJSValue()

	oldValue.Call("replaceWith", newValue)
	n.current = &newValue
}
