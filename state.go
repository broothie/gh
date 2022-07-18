package jog

type State struct {
	parent        *State
	values        map[string]any
	subscriptions map[string][]subscription
}

type subscription struct {
	node    *Node
	watcher StateWatcher
}

func NewState() *State {
	return &State{
		values:        make(map[string]any),
		subscriptions: make(map[string][]subscription),
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
				node := subscription.watcher(value)
				subscription.node.update(node)
				subscription.node = node
			}

			return
		}
	}

	s.values[name] = value
}

type StateWatcher func(value any) *Node

func (s *State) Watch(name string, watcher StateWatcher) *Node {
	node := watcher(s.Get(name))

	for current := s; current != nil; current = current.parent {
		if _, found := current.values[name]; found {
			current.subscriptions[name] = append(current.subscriptions[name], subscription{
				node:    node,
				watcher: watcher,
			})

			break
		}
	}

	return node
}

func (s *State) NewChild() *State {
	return &State{
		parent:        s,
		values:        make(map[string]any),
		subscriptions: make(map[string][]subscription),
	}
}
