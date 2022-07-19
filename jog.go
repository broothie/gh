package jog

func Mount(id string, builder BuildFunc) {
	element := document.Call("getElementById", id)
	element.Call("appendChild", builder(NewState()).Generate())
}

func Wait() {
	select {}
}
