package block

import (
	"bytes"
	"context"
	"encoding/json"
)

func init() {
	fullBlockStream = make(chan WithTxs)
}

// WithTxs is used to couple blocks with txs data together for our map
type WithTxs struct {
	B   *Block
	Txs []TxStatus
}

// Full returns true if all txs are matched inside this block
func (w WithTxs) Full() bool {
	// if we don't have yet received the block itself, exit here so we don't panic
	if w.B == nil {
		return false
	}

	// helper function telling that hash is present in our txStatus slice
	present := func(hash []byte, txs []TxStatus) bool {
		for _, tx := range txs {
			if bytes.Equal(tx.Hash, hash) {
				return true
			}
		}
		return false
	}

	for _, tx := range w.B.Txs {
		if !present(tx.Hash(), w.Txs) {
			return false
		}
	}

	return true
}

// JSON marshall struct to json format
func (w WithTxs) JSON() []byte {
	bytes, _ := json.Marshal(w)
	return bytes
}

// fullBlockStream represents a channel to which blocks are getting send as they are pushed to our map
// it is used inside block_map.go internally and must always be initialized
var fullBlockStream chan WithTxs

// initiliazed stream is a read-only stream which we return to user, it wraps around
// internal fullBlockStream and uses orDone pattern in it's getter function for convenient use
var initializedStream chan WithTxs

// GetBlockStream returns read-only channel of completed blocks
func GetBlockStream(ctx context.Context, tendermintRPC string) (<-chan WithTxs, error) {
	// we already have our block stream created, with workers pushing blocks and txs to it
	if initializedStream != nil {
		return initializedStream, nil
	}

	blocks, txs, err := Subscribe(ctx, tendermintRPC)
	// couldn't connect to node
	if err != nil {
		return nil, err
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
	initializedStream = make(chan WithTxs)
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

	return initializedStream, nil
}
