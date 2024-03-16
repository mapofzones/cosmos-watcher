package cosmos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cometbft/cometbft/rpc/client/http"
	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
)

// GetBlock queries tendermint rpc at provided height and formats block
// it does that by fetching the block itself and then querying each tx in the block
func GetBlock(ctx context.Context, client *http.HTTP, N int64) (block.Block, error) {
	Block, err := client.Block(ctx, &N)
	if err != nil {
		return block.Block{}, err
	}

	s := []block.TxStatus{}
	txRes := []*abci.ResponseDeliverTx{}
	results, err := client.BlockResults(ctx, &N)
	if err != nil {
		for _, tx := range Block.Block.Txs {
			res, err := client.Tx(ctx, tx.Hash(), false)
			if errors.Is(err, errors.New("Tx")) {
				return block.Block{}, fmt.Errorf("Transaction does not exist: %w", err)
			} else if err != nil {
				return block.Block{}, err
			}
			s = append(s, block.TxStatus{
				ResultCode: res.TxResult.Code,
				Hash:       tx.Hash(),
				Height:     res.Height,
			})
			txRes = append(txRes, &res.TxResult)
		}
	} else {
		if len(results.TxsResults) != len(Block.Block.Txs) {
			return block.Block{}, fmt.Errorf("wrong number of block txs & block txs result: %w", err)
		} else {
			for i := 0; i < len(results.TxsResults); i++ {
				s = append(s, block.TxStatus{
					ResultCode: results.TxsResults[i].Code,
					Hash:       Block.Block.Txs[i].Hash(),
					Height:     N,
				})
				txRes = append(txRes, results.TxsResults[i])
			}
		}

	}

	return block.Block{
		ChainID:      Block.Block.ChainID,
		Height:       Block.Block.Height,
		T:            Block.Block.Time,
		Txs:          Block.Block.Txs,
		Results:      s,
		BlockResults: results,
		TxsResults:   txRes,
	}, nil
}

// block range queries tendermint node for blocks from [first,last]
// callers need to check if channel return by the function to know that
// it had finished working or some error occurred
// the function does not guarantee that all blocks will be fetched, so callers
// should also keep track of that
func BlockRange(ctx context.Context, client *http.HTTP, first, last int64) <-chan block.Block {
	blockStream := make(chan block.Block)

	go func() {
		defer close(blockStream)
		duration := time.Second * 1500
		timer := time.NewTimer(duration)
		for N := first; N <= last; N++ {
			block, err := GetBlock(ctx, client, N)
			if err != nil {
				log.Println(err)
				return
			}
			select {
			case blockStream <- block:
				timer.Reset(duration)
			case <-timer.C:
				log.Println("Timer finished exit")
				os.Exit(1)
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockStream
}
