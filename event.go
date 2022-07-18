package jog

import "syscall/js"

type Event struct {
	JSValue js.Value
}

type EventListener interface {
	HandleEvent(event Event) any
}

type Listener func(event Event) any

func (f Listener) HandleEvent(event Event) any {
	return f(event)
}
