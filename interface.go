package watcher

import (
	"context"
	"errors"

	watcher "github.com/mapofzones/cosmos-watcher/types"
	tm "github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit"
)

// Watcher interface listens on the listenAddr and sends data to the sendAddr
type Watcher interface {
	Watch(context.Context) error
}

// NewWatcher returns Watcher implementation based on implementation specified by t and logger l
func NewWatcher(t watcher.WType, tendermintRPC, rabbitMQAddr string) (Watcher, error) {
	switch t {
	case watcher.TmRabbit:
		return tm.NewWatcher(tendermintRPC, rabbitMQAddr)
	default:
		return nil, errors.New("invalid type argument")
	}

}
