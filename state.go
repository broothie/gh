package gh

import "syscall/js"

type State struct {
	parent        *State
	values        map[string]any
	subscriptions map[string][]*subscription
}

type subscription struct {
	current js.Value
	watcher StateWatcher
}

func NewState() *State {
	return &State{
		values:        make(map[string]any),
		subscriptions: make(map[string][]*subscription),
	}
}

func (s *State) Get(name string) any {
	if value, found := s.values[name]; found {
		return value
	}

	if s.parent != nil {
		return s.parent.Get(name)
	}

	return nil
}

func (s *State) With(values map[string]any) {
	for name, value := range values {
		s.Set(name, value)
	}
}

func (s *State) Set(name string, value any) {
	for current := s; current != nil; current = current.parent {
		if _, found := current.values[name]; found {
			current.values[name] = value

			for _, subscription := range current.subscriptions[name] {
				if !subscription.current.Get("isConnected").Truthy() {
					continue
				}

				jsValue := subscription.watcher(value).Generate()
				subscription.current.Call("replaceWith", jsValue)
				subscription.current = jsValue
			}

			return
		}
	}

	s.values[name] = value
}

type StateWatcher func(value any) Generator

func (s *State) Watch(name string, watcher StateWatcher) GeneratorFunc {
	sub := &subscription{watcher: watcher}

	for current := s; current != nil; current = current.parent {
		if _, found := current.values[name]; found {
			current.subscriptions[name] = append(current.subscriptions[name], sub)
			break
		}
	}

	return func() js.Value {
		value := watcher(s.Get(name)).Generate()
		sub.current = value
		return value
	}
}

func (s *State) NewChild() *State {
	return &State{
		parent:        s,
		values:        make(map[string]any),
		subscriptions: make(map[string][]*subscription),
	}
}
