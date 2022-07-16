package jog

import (
	"context"
	"syscall/js"
)

type Node interface {
	JSValue(ctx context.Context) js.Value
}
