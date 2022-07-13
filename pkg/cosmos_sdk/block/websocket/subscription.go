package cosmos

import (
	"context"
	"log"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

// Subscribe dials tendermint rpc and returns two streams, one for committed blocks, one for transactions that occurred
func Subscribe(ctx context.Context, client *http.HTTP) (<-chan block.TmBlock, <-chan block.TxStatus, error) {
	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 1")
	blockChan, err := client.Subscribe(ctx, "", "tm.event = 'NewBlock'", 10000)
	if err != nil {
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 2")
		return nil, nil, err
	}

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 3")
	txChan, err := client.Subscribe(ctx, "", "tm.event = 'Tx'", 10000)
	if err != nil {
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 4")
		return nil, nil, err
	}

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 5")
	return toBlockStream(ctx, blockChan), toTxStream(ctx, txChan), nil
}

// this function returns values from c channel but also checks if it's closed and respects context
func toBlockStream(ctx context.Context, c <-chan coretypes.ResultEvent) <-chan block.TmBlock {
	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 6")
	blockStream := make(chan block.TmBlock)

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 7")
	go func() {
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 8")
		defer close(blockStream)
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 9")
		for {
			log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 10")
			select {
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 11")
				return
			case v, ok := <-c:
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 12")
				if !ok {
					return
				}
				select {
				case blockStream <- block.BlockFromTmResultBlock(v.Data.(types.EventDataNewBlock)):
				case <-ctx.Done():
				}
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 13")
			}
			log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 14")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 15")
	return blockStream
}

// this function returns values from c channel but also checks if it's closed and respects context
func toTxStream(ctx context.Context, c <-chan coretypes.ResultEvent) <-chan block.TxStatus {
	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 16")
	txStream := make(chan block.TxStatus)

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 17")
	go func() {
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 18")
		defer close(txStream)
		log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 19")
		for {
			log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 20")
			select {
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 21")
				return
			case v, ok := <-c:
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 22")
				if !ok {
					return
				}
				select {
				case txStream <- block.TxStatusFromTmResultTx(v.Data.(types.EventDataTx)):
				case <-ctx.Done():
				}
				log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 23")
			}
			log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 24")
		}
	}()

	log.Println("pkg.cosmos_sdk.block.websocket.subscription.go - 25")
	return txStream
}
