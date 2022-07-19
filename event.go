package gh

import "syscall/js"

type EventListener interface {
	HandleEvent(event Event) any
}

type Listener func(event Event) any

func (f Listener) HandleEvent(event Event) any {
	return f(event)
}

type Event struct {
	js.Value
}

func (e Event) Target() Target {
	return Target{JSValue: e.Get("target")}
}

type Target struct {
	JSValue js.Value
}

func (t Target) Value() js.Value {
	return t.JSValue.Get("value")
}
