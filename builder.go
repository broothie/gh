package jog

type Builder interface {
	Build(state *State) *Node
}

type BuilderFunc func(state *State) *Node

func (b BuilderFunc) Build(state *State) *Node {
	return b(state)
}
