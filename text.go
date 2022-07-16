package jog

import (
	"context"
	"html"
	"syscall/js"
)

type Text string

func (t Text) JSValue(context.Context) js.Value {
	return Document().Call("createTextNode", html.EscapeString(string(t)))
}
