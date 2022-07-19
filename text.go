package gh

import (
	"fmt"
	"syscall/js"
)

func Text(value any) GeneratorFunc {
	return func() js.Value {
		return Document.CreateTextNode(fmt.Sprint(value))
	}
}
