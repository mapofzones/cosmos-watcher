package cosmos

import (
	"context"
	"errors"
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"
	"log"
	"os"
	"time"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	"github.com/tendermint/tendermint/rpc/client/http"
)

// GetBlock queries tendermint rpc at provided height and formats block
// it does that by fetching the block itself and then querying each tx in the block
func GetBlock(ctx context.Context, client *http.HTTP, N int64) (block.Block, error) {
	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 1")
	Block, err := client.Block(ctx, &N)
	if err != nil {
		return block.Block{}, err
	}

	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 2")
	s := []block.TxStatus{}
	txRes := []*abci.ResponseDeliverTx{}
	results, err := client.BlockResults(ctx, &N)
	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 3")
	if err != nil {
		log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 4")
		for _, tx := range Block.Block.Txs {
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 5")
			res, err := client.Tx(ctx, tx.Hash(), false)
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 6")
			if errors.Is(err, errors.New("Tx")) {
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 7")
				return block.Block{}, fmt.Errorf("Transaction does not exist: %w", err)
			} else if err != nil {
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 8")
				return block.Block{}, err
			}
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 9")
			s = append(s, block.TxStatus{
				ResultCode: res.TxResult.Code,
				Hash:       tx.Hash(),
				Height:     res.Height,
			})
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 10")
			txRes = append(txRes, &res.TxResult)
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 11")
		}
	} else {
		log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 12")
		if len(results.TxsResults) != len(Block.Block.Txs) {
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 13")
			return block.Block{}, fmt.Errorf("wrong number of block txs & block txs result: %w", err)
		} else {
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 14")
			for i := 0; i < len(results.TxsResults); i++ {
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 15")
				s = append(s, block.TxStatus{
					ResultCode: results.TxsResults[i].Code,
					Hash:       Block.Block.Txs[i].Hash(),
					Height:     N,
				})
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 16")
				txRes = append(txRes, results.TxsResults[i])
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 17")
			}
		}

	}

	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 18")
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
	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 19")
	blockStream := make(chan block.Block)

	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 20")
	go func() {
		log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 21")
		defer close(blockStream)
		duration := time.Second * 150
		timer := time.NewTimer(duration)
		log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 22")
		for N := first; N <= last; N++ {
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 23")
			block, err := GetBlock(ctx, client, N)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 24")
			select {
			case blockStream <- block:
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 25")
				timer.Reset(duration)
			case <-timer.C:
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 26")
				log.Println("Timer finished exit")
				os.Exit(1)
			case <-ctx.Done():
				log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 27")
				return
			}
			log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 28")
		}
		log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 29")
	}()

	log.Println("pkg.cosmos_sdk.block.crawler.crawler.go - 30")
	return blockStream
}
