package jog

import "syscall/js"

var (
	document = js.Global().Get("document")
	head     = js.Global().Get("head")
	body     = js.Global().Get("body")
	Console  = console(js.Global().Get("console"))
)

type console js.Value

func (c console) Log(v ...any) {
	js.Value(c).Call("log", v...)
}
