package jog

import "syscall/js"

type Text string

func (t Text) Generate() js.Value {
	return document.Call("createTextNode", string(t))
}
