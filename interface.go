package watcher

import (
	"errors"

	"github.com/attractor-spectrum/cosmos-watcher/tx"
	config "github.com/attractor-spectrum/cosmos-watcher/x/config"
	tm "github.com/attractor-spectrum/cosmos-watcher/x/tendermint-rabbit"
)

// WType is used to choose Watcher implementation
type WType = uint32

const (
	// TmRabbit is tendermint and rabbitMQ watcher interface implementation
	TmRabbit WType = iota
)

// Tx is non-blockchain-specific transaction representation
type Tx = tx.Tx

// Tx type aliases
type Type = tx.Type

const (
	Transfer   = tx.Transfer
	Stake      = tx.Stake
	Unstake    = tx.Unstake
	IbcSend    = tx.IbcSend
	IbcRecieve = tx.IbcRecieve
	Other      = tx.Other
)

// Watcher interface listens on the listenAddr and sends data to the sendAddr
type Watcher interface {
	Watch() error
}

// NewWatcher returns Watcher implementation based on implementation specified by t and logger l
func NewWatcher(t WType, config *config.Config) (Watcher, error) {
	switch t {
	case TmRabbit:
		return tm.NewWatcher(config)
	default:
		return nil, errors.New("invalid type argument")
	}

}
