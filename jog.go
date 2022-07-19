package gh

type BuildFunc func(state *State) Generator

func Mount(id string, builder BuildFunc) {
	element := Document.Call("getElementById", id)
	element.Call("appendChild", builder(NewState()).Generate())
}

func Wait() {
	select {}
}
