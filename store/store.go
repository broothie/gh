package store

import (
	"syscall/js"

	"github.com/broothie/gh"
	"github.com/broothie/gh/util"
	"github.com/samber/lo"
)

type Store[T any] struct {
	value         T
	subscriptions []*subscription[T]
}

func New[T any](value T) *Store[T] {
	return &Store[T]{value: value}
}

func (s *Store[T]) Get() T {
	return util.DeepCopy(s.value)
}

func (s *Store[T]) Modify(modify func(*T)) {
	cp := util.DeepCopy(s.value)
	modify(&cp)
	s.Update(cp)
}

func (s *Store[T]) Update(value T) {
	previous := s.value
	s.value = value

	var removals []*subscription[T]
	for _, sub := range s.subscriptions {
		if !sub.current.Call("isConnected").Truthy() {
			removals = append(removals, sub)
			continue
		}

		if sub.changed(util.DeepCopy(previous), util.DeepCopy(value)) {
			go func(sub *subscription[T]) {
				jsValue := sub.handler(util.DeepCopy(value)).Generate()
				sub.current.Call("replaceWith", jsValue)
				sub.current = jsValue
			}(sub)
		}
	}

	s.subscriptions = lo.Without(s.subscriptions, removals...)
}

type Changed[T any] func(T, T) bool
type Handler[T any] func(T) gh.Generator

type subscription[T any] struct {
	current js.Value
	changed Changed[T]
	handler Handler[T]
}

func (s *Store[T]) Watch(changed Changed[T], handler Handler[T]) gh.Generator {
	return gh.GeneratorFunc(func() js.Value {
		value := handler(util.DeepCopy(s.value)).Generate()

		s.subscriptions = append(s.subscriptions, &subscription[T]{
			current: value,
			changed: changed,
			handler: handler,
		})

		return value
	})
}

func WatchMapLengthChanged[K comparable, V any](a, b map[K]V) bool {
	return len(a) != len(b)
}

func WatchSliceLengthChanged[V any](a, b []V) bool {
	return len(a) != len(b)
}

func WatchSelection[T any, U comparable](store *Store[T], selector func(T) U, handler func(U) gh.Generator) gh.Generator {
	return store.Watch(
		func(a, b T) bool { return selector(a) != selector(b) },
		func(value T) gh.Generator { return handler(selector(value)) },
	)
}
