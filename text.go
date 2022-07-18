package jog

import "syscall/js"

func Text(s string) *Node {
	return &Node{generator: text(s)}
}

type text string

func (t text) Generate() js.Value {
	return document.Call("createTextNode", string(t))
}
