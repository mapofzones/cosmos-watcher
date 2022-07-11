package cosmos

import (
	"context"
	"log"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func init() {
	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 1")
	fullBlockStream = make(chan block.WithTxs)
	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 2")
}

// fullBlockStream represents a channel to which blocks are getting send as they are pushed to our map
// it is used inside block_map.go internally and must always be initialized
var fullBlockStream chan block.WithTxs

// initiliazed stream is a read-only stream which we return to user, it wraps around
// internal fullBlockStream and uses orDone pattern in it's getter function for convenient use
var initializedStream chan block.WithTxs

// BlockStream returns read-only channel of completed blocks
func BlockStream(ctx context.Context, c *http.HTTP) <-chan block.WithTxs {
	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 3")
	// we already have our block stream created, with workers pushing blocks and txs to it
	if initializedStream != nil {
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 4")
		return initializedStream
	}

	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 5")
	blocks, txs, err := Subscribe(ctx, c)
	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 6")
	// couldn't connect to node
	if err != nil {
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 7")
		log.Println(err)
		return nil
	}

	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 8")
	// process each block
	go func() {
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 9")
		for block := range blocks {
			pushBlock(block)
		}
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 10")
	}()

	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 11")
	// process each tx
	go func() {
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 12")
		for tx := range txs {
			pushTx(tx)
		}
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 13")
	}()

	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 14")
	// orDone pattern so we don't have to check for closed channel
	initializedStream = make(chan block.WithTxs)
	go func() {
		log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 15")
		defer close(initializedStream)
		for {
			select {
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 16")
				return
			case block, ok := <-fullBlockStream:
				log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 17")
				if ok {
					log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 18")
					select {
					case initializedStream <- block:
					case <-ctx.Done():
					}
					log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 19")
				}
				log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 20")
			}
			log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 21")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.websocket.block_stream.go - 22")
	return initializedStream
}
