package watcher

import "time"

// Message represents some event / action
// that happened inside of blockchain
type Message interface {
	Type() string
}

// Block is an interface providing array of messages (state transitions)
// as well as helper methods providing usefull information about the block
// all Block interface implementations are assumed to return
// messages in order in which those events / messages happened
type Block interface {
	Height() int64
	ChainID() string
	Time() time.Time
	Messages() []Message
}
