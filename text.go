package gh

import "syscall/js"

type Text string

func (t Text) Generate() js.Value {
	return Document.CreateTextNode(string(t))
}
