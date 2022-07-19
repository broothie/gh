package promise

type status string

const (
	statusPending   status = "pending"
	statusFulfilled status = "fulfilled"
	statusRejected  status = "rejected"
)

type Promise[T any] struct {
	resolve func(T)
	reject  func(error)

	status status
	result T
	error  error
}

func (p *Promise[T]) Then(resolve func(T) ) {
	if p.status == statusFulfilled {
		resolve(p.result)
	} else {
		p.resolve = resolve
	}
}

func (p *Promise[T]) Catch(reject func(error)) {
	if p.status == statusRejected {
		reject(p.error)
	} else {
		p.reject = reject
	}
}

func From[T any](f func() (T, error)) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		if result, err := f(); err != nil {
			reject(err)
		} else {
			resolve(result)
		}
	})
}

func New[T any](f func(resolve func(T), reject func(error))) *Promise[T] {
	promise := &Promise[T]{status: statusPending}

	go f(
		func(result T) {
			promise.result = result
			promise.status = statusFulfilled

			if promise.resolve != nil {
				promise.resolve(result)
			}
		},
		func(err error) {
			promise.error = err
			promise.status = statusRejected

			if promise.reject != nil {
				promise.reject(err)
			}
		},
	)

	return promise
}
