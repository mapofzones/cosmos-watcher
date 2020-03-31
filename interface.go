package watcher

import (
	"errors"
	"log"
	"os"

	tm "github.com/attractor-spectrum/cosmos-watcher/x/tendermint-rabbit"
)

// WType is used to choose Watcher implementation
type WType = uint32

const (
	// TmRabbit is tendermint and rabbitMQ watcher interface implementation
	TmRabbit WType = iota
)

// Watcher interface listens on the listenAddr and sends data to the sendAddr
type Watcher interface {
	Watch() error
}

// NewWatcher returns Watcher implementation based on implementation specified by t and logger l
func NewWatcher(t WType, l *log.Logger) (Watcher, error) {
	if l == nil {
		l = log.New(os.Stdout, "", log.LstdFlags)
	}
	switch t {
	case TmRabbit:
		return tm.NewWatcher(l)
	default:
		return nil, errors.New("invalid type argument")
	}

}
