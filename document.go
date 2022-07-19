package gh

import "syscall/js"

var Document = document{Value: js.Global().Get("document")}

type document struct {
	js.Value
}

func (d document) CreateElement(name string) js.Value {
	return d.Call("createElement", name)
}

func (d document) CreateTextNode(text string) js.Value {
	return d.Call("createTextNode", text)
}

func (d document) AppendChild(child js.Value) js.Value {
	return d.Call("appendChild", child)
}
