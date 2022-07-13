package watcher

import (
	"context"
	"errors"
	"log"
	"os"
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
	log.Println("pkg.watcher.go - 1")
	return &Watcher{
		BlockStream: blockStream,
		Queue:       rabbitQueue,
	}
}

// WatchWithTimeout is used to receive blocks from Block Stream
// and send them to Queue
func (w *Watcher) WatchWithTimeout(ctx context.Context, timeout time.Duration) error {
	log.Println("pkg.watcher.go - 2")
	duration := time.Second * 150
	timer := time.NewTimer(duration)
	for {
		log.Println("pkg.watcher.go - 3")
		select {
		// get block
		case block, ok := <-w.BlockStream:
			log.Println("pkg.watcher.go - 4")
			timer.Reset(duration)
			if !ok {
				log.Println("pkg.watcher.go - 5")
				return errors.New("block channel is closed")
			}
			log.Println("pkg.watcher.go - 6")
			select {
			// send it for further processing
			case w.Queue <- block:
			case <-ctx.Done():
			}
			log.Println("pkg.watcher.go - 7")
		// timeout if we did not receive any data
		case <-time.Tick(timeout):
			log.Println("pkg.watcher.go - 8")
			return errors.New("timeout: did not receive any blocks for 10 minutes")
		case <-timer.C:
			log.Println("pkg.watcher.go - 9")
			log.Println("Timer finished exit")
			os.Exit(1)
		case <-ctx.Done():
			log.Println("pkg.watcher.go - 10")
			return nil
		}
		log.Println("pkg.watcher.go - 11")
	}
}
