package jog

func Mount(id string, builder Builder) {
	element := document.Call("getElementById", id)
	element.Call("appendChild", builder.Build(NewState()).ToJSValue())
}

func Wait() {
	select {}
}
