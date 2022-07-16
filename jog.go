package jog

import (
	"context"
	"syscall/js"
)

func Mount(id string, node Node) {
	element := Document().Call("getElementById", id)
	// element.Call("appendChild", node.JSValue(context.TODO()))

	element.Call("appendChild", view().JSValue(newContextWithJogCtx()))
}

func Wait() {
	<-make(chan struct{})
}

type Func func(ctx context.Context) Node

func (f Func) JSValue(ctx context.Context) js.Value {
	return f(ctx).JSValue(ctx)
}

func view() Func {
	return func(ctx context.Context) Node {
		p := P(nil)

		p.subscriptions = append(p.subscriptions, "query")

		return Div(nil,
			Input(Attr{"type": "text", "onkeyup": js.FuncOf(func(_ js.Value, args []js.Value) any {
				query := args[0].Get("target").Get("value")
				jCtx := jogCtx(ctx)
				p := jCtx.subscriptions["query"][0]
				newP := P(nil, Text(query.String())).JSValue(ctx)

				p.Call("replaceWith", newP)
				jCtx.subscriptions["query"][0] = newP

				return nil
			})}),
			p,
		)
	}
}

var jogContextKey struct{}

type jogContext struct {
	subscriptions map[string][]js.Value
}

func jogCtx(ctx context.Context) *jogContext {
	return ctx.Value(jogContextKey).(*jogContext)
}

func newContextWithJogCtx() context.Context {
	return context.WithValue(context.Background(), jogContextKey, &jogContext{subscriptions: make(map[string][]js.Value)})
}
