package watcher

import "time"

// StateTransition represents some event / action
// that happened inside of blockchain
type StateTransition interface {
	Type() string
}

// Block is an interface providing array of state transitions
// as well as helper methods providing usefull information about the block
// all Block interface implementations are assumed to return
// state transitions in order in which those transitions happened
type Block interface {
	Height() int64
	ChainID() string
	Time() time.Time
	Transitions() []StateTransition
}
