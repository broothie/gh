package gh

func Mount(id string, generator Generator) {
	element := Document.Call("getElementById", id)
	element.Call("appendChild", generator.Generate())
}

func Wait() {
	select {}
}
