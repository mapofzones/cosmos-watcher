package listener

import (
	"errors"
	"log"
	"os"

	tm "github.com/attractor-spectrum/cosmos-watcher/listener/x/tendermint-rabbit"
)

// LType is used to choose Listener implementation
type LType = uint32

const (
	// TmRabbit is tendermint and rabbitMQ listener interface implementation
	TmRabbit LType = iota
)

// Listener interface listens on the listenAddr and sends data to the sendAddr
type Listener interface {
	ListenAndServe() error
}

// NewListener returns Listener implementation based on implementation specified by t and logger l
func NewListener(t LType, l *log.Logger) (Listener, error) {
	if l == nil {
		l = log.New(os.Stdout, "", log.LstdFlags)
	}
	switch t {
	case TmRabbit:
		return tm.NewListener(l)
	default:
		return nil, errors.New("invalid type argument")
	}

}
