package watcher

type StateTransition interface {
	Type() string
}

type Block interface {
	Height() int64
	ChainID() string
	Transitions() []StateTransition
}
