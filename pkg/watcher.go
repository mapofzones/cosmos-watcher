package watcher

import (
	"context"
	"errors"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
)

// Watcher is a wrapper around two channels:
// first one (block stream) for getting and preprocessing blocks
// second (Queue) for sending them somewhere else for further processing
type Watcher struct {
	BlockStream <-chan watcher.Block
	Queue       chan<- watcher.Block
}

// NewWatcher returns instanciated Watcher
func NewWatcher(ctx context.Context, blockStream <-chan watcher.Block, rabbitQueue chan<- watcher.Block) *Watcher {
	return &Watcher{
		BlockStream: blockStream,
		Queue:       rabbitQueue,
	}
}

// WatchWithTimeout is used to receive blocks from Block Stream
// and send them to Queue
func (w *Watcher) WatchWithTimeout(ctx context.Context, timeout time.Duration) error {
	for {
		select {
		// get block
		case block, ok := <-w.BlockStream:
			if !ok {
				return errors.New("block channel is closed")
			}
			select {
			// send it for further processing
			case w.Queue <- block:
			case <-ctx.Done():
			}
		// timeout if we did not receive any data
		case <-time.Tick(timeout):
			return errors.New("timeout: did not receive any blocks for 10 minutes")
		case <-ctx.Done():
			return nil
		}
	}
}
