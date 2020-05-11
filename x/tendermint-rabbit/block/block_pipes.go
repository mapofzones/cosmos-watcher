package block

import (
	"context"
	"log"

	crawler "github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block/crawler"
	block "github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block/types"
	websocket "github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block/websocket"
	"github.com/tendermint/tendermint/rpc/client/http"
)

// normalizedStream is used to format websocket data correspondingly to our exported block structure
func normalizedStream(ctx context.Context, blockWithTxs <-chan block.WithTxs) <-chan block.Block {
	blockStream := make(chan block.Block)

	go func() {
		defer close(blockStream)
		for {
			select {
			case Block, ok := <-blockWithTxs:
				if !ok {
					return
				}
				select {
				case blockStream <- block.Normalize(Block):
				case <-ctx.Done():
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockStream
}

func ordered(ctx context.Context, stream <-chan block.Block, startHeight int64) <-chan block.Block {
	orderedStream := make(chan block.Block)

	go func() {
		defer close(orderedStream)
		expectedHeight := startHeight

		blocks := map[int64]block.Block{}
		for {
			// go through the local cache
			if block, ok := blocks[expectedHeight]; ok {
				select {
				case orderedStream <- block:
					delete(blocks, expectedHeight)
					expectedHeight++
					continue
				case <-ctx.Done():
					return
				}
			}
			break
		}
		select {
		case block, ok := <-orderedStream:
			// return since the source channel closed
			if !ok {
				return
			}

			// should not happen, but worth checking in case of a bug somewhere in the pipe
			if block.Height < expectedHeight {
				break
			}

			// if we got the block we expected to get
			if block.Height == expectedHeight {
				select {
				case orderedStream <- block:
					expectedHeight++
				case <-ctx.Done():
					return
				}
				// if this is not the block we need
			} else {
				// put it in local cache
				blocks[block.Height] = block
			}
		case <-ctx.Done():
			return
		}
	}()

	return orderedStream
}

func BlockStream(ctx context.Context, client *http.HTTP, startHeight int64) <-chan block.Block {
	blockStream := make(chan block.Block)

	go func() {
		defer close(blockStream)
		status, err := client.Status()
		if err != nil {
			log.Println(err)
			return
		}
		currentHeight := status.SyncInfo.LatestBlockHeight
		// if we have a lot of catching up to do
		for currentHeight-startHeight > 10 {
			blocks := crawler.BlockRange(ctx, client, startHeight, currentHeight)
			for block := range blocks {
				select {
				case blockStream <- block:
					startHeight++
				case <-ctx.Done():
					return
				}
			}

			// update current latest blockchain height to know if we need another iteration of this loop
			status, err := client.Status()
			if err != nil {
				log.Println(err)
				return
			}
			currentHeight = status.SyncInfo.LatestBlockHeight
		}

		// alias variable to make stuff less confusing
		// since we synced up, start height == lastProcessedBlock
		lastProcessedBlock := startHeight

		// if we reached this point in the function, we can start listening on the websocket in realtime
		websocketStream := normalizedStream(ctx, websocket.BlockStream(ctx, client))

		// get one block from websocket to know how many blocks we have missed while switching from
		// crawler to websocket block stream
		var pivotBlock block.Block
		select {
		case block, ok := <-websocketStream:
			if !ok {
				return
			}
			pivotBlock = block
		case <-ctx.Done():
			return
		}

		// last time we catch up with node
		for block := range crawler.BlockRange(ctx, client, lastProcessedBlock, pivotBlock.Height) {
			select {
			case blockStream <- block:
			case <-ctx.Done():
				return
			}
		}

		// send pivot block
		select {
		case blockStream <- pivotBlock:
		case <-ctx.Done():
			return
		}

		// switch to websocket at last
		for {
			select {
			case blockStream <- <-websocketStream:
			case <-ctx.Done():
				return
			}
		}

	}()

	return blockStream
}
