package block

import (
	"context"

	block "github.com/mapofzones/cosmos-watcher/pkg/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

// Subscribe dials tendermint rpc and returns two streams, one for committed blocks, one for transactions that occurred
func Subscribe(ctx context.Context, client *http.HTTP) (<-chan block.TmBlock, <-chan block.TxStatus, error) {
	blockChan, err := client.Subscribe(context.Background(), "", "tm.event = 'NewBlock'", 10000)
	if err != nil {
		return nil, nil, err
	}

	txChan, err := client.Subscribe(context.Background(), "", "tm.event = 'Tx'", 10000)
	if err != nil {
		return nil, nil, err
	}

	return toBlockStream(ctx, blockChan), toTxStream(ctx, txChan), nil
}

// this function returns values from c channel but also checks if it's closed and respects context
func toBlockStream(ctx context.Context, c <-chan coretypes.ResultEvent) <-chan block.TmBlock {
	blockStream := make(chan block.TmBlock)

	go func() {
		defer close(blockStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case blockStream <- block.BlockFromTmResultBlock(v.Data.(types.EventDataNewBlock)):
				case <-ctx.Done():
				}
			}
		}
	}()

	return blockStream
}

// this function returns values from c channel but also checks if it's closed and respects context
func toTxStream(ctx context.Context, c <-chan coretypes.ResultEvent) <-chan block.TxStatus {
	txStream := make(chan block.TxStatus)

	go func() {
		defer close(txStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case txStream <- block.TxStatusFromTmResultTx(v.Data.(types.EventDataTx)):
				case <-ctx.Done():
				}
			}
		}
	}()

	return txStream
}
