package cosmos

import (
	"context"
	"errors"
	"fmt"
	"log"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
)

// GetBlock queries tendermint rpc at provided height and formats block
// it does that by fetching the block itself and then querying each tx in the block
func GetBlock(client *http.HTTP, N int64) (block.Block, error) {
	Block, err := client.Block(&N)
	if err != nil {
		return block.Block{}, err
	}

	s := []block.TxStatus{}
	for _, tx := range Block.Block.Txs {
		res, err := client.Tx(tx.Hash(), false)
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
	}

	return block.Block{
		ChainID: Block.Block.ChainID,
		Height:  Block.Block.Height,
		T:       Block.Block.Time,
		Txs:     Block.Block.Txs,
		Results: s,
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
		for N := first; N <= last; N++ {
			block, err := GetBlock(client, N)
			if err != nil {
				log.Println(err)
				return
			}
			select {
			case blockStream <- block:
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockStream
}
