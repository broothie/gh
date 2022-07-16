package jog

import "syscall/js"

func Console() js.Value {
	return js.Global().Get("console")
}

func ConsoleLog(messages ...any) {
	Console().Call("log", messages...)
}

func Document() js.Value {
	return js.Global().Get("document")
}

func Head() js.Value {
	return Document().Get("head")
}

func Body() js.Value {
	return Document().Get("body")
}
