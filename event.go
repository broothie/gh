package jog

import "syscall/js"

type Event struct {
	js.Value
}

type EventListener interface {
	HandleEvent(event Event)
}

type EventListenerFunc func(event Event)

func (f EventListenerFunc) HandleEvent(event Event) {
	f(event)
}
