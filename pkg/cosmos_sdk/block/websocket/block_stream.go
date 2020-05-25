package cosmos

import (
	"context"
	"log"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func init() {
	fullBlockStream = make(chan block.WithTxs)
}

// fullBlockStream represents a channel to which blocks are getting send as they are pushed to our map
// it is used inside block_map.go internally and must always be initialized
var fullBlockStream chan block.WithTxs

// initiliazed stream is a read-only stream which we return to user, it wraps around
// internal fullBlockStream and uses orDone pattern in it's getter function for convenient use
var initializedStream chan block.WithTxs

// BlockStream returns read-only channel of completed blocks
func BlockStream(ctx context.Context, c *http.HTTP) <-chan block.WithTxs {
	// we already have our block stream created, with workers pushing blocks and txs to it
	if initializedStream != nil {
		return initializedStream
	}

	blocks, txs, err := Subscribe(ctx, c)
	// couldn't connect to node
	if err != nil {
		log.Println(err)
		return nil
	}

	// process each block
	go func() {
		for block := range blocks {
			pushBlock(block)
		}
	}()

	// process each tx
	go func() {
		for tx := range txs {
			pushTx(tx)
		}
	}()

	// orDone pattern so we don't have to check for closed channel
	initializedStream = make(chan block.WithTxs)
	go func() {
		defer close(initializedStream)
		for {
			select {
			case <-ctx.Done():
				return
			case block, ok := <-fullBlockStream:
				if ok {
					select {
					case initializedStream <- block:
					case <-ctx.Done():
					}
				}
			}
		}
	}()

	return initializedStream
}
