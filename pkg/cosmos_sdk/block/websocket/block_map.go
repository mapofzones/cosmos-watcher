package cosmos

import (
	"log"
	"sync"

	block "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
)

func init() {
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 1")
	blocks = make(map[int64]block.WithTxs)
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 2")
	lock = &sync.Mutex{}
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 3")
}

// assume that each watcher is launched in separate binary
// so all records in map belong to one chain

// blocks is a global map object where tx data is stored
// should not be used directly, instead use helper methods in the package
var blocks map[int64]block.WithTxs
var lock *sync.Mutex

// pushBlock puts block in map and sends its over a full block channel for further parsing
func pushBlock(b block.TmBlock) {
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 4")
	// check if block has any transactions
	if len(b.Txs) == 0 {
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 5")
		// send block for further processing
		go func() { fullBlockStream <- block.WithTxs{B: &b, Txs: []block.TxStatus{}} }()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 6")
		return
	}

	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 7")
	// write to map critical section
	lock.Lock()
	// if WithTx structure is already present in the map
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 8")
	if _, ok := blocks[b.Height]; ok {
		blocks[b.Height] = block.WithTxs{B: &b, Txs: blocks[b.Height].Txs}
	} else {
		blocks[b.Height] = block.WithTxs{B: &b, Txs: make([]block.TxStatus, len(b.Txs))}
	}

	// if we have filled the block with all necessary tx codes, send it to block stream and delete it from map
	// we need local copy because we might delete it earlier that block is getting send to stream
	localCopy := blocks[b.Height]

	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 9")
	// we unlock here because full function could take some time to complete
	lock.Unlock()
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 10")
	if localCopy.Full() {
		// send data to channel
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 11")
		go func() { fullBlockStream <- localCopy }()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 12")
		// delete data which we sent to channel already
		lock.Lock()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 13")
		delete(blocks, b.Height)
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 14")
		lock.Unlock()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 15")
		return
	}
}

// pushTx puts transaction in our map and, if block if full, sends it to block channel
func pushTx(tx block.TxStatus) {
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 16")
	// our write to map critical section
	lock.Lock()
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 17")
	// if WithTx structure is already present in the map
	if _, ok := blocks[tx.Height]; ok {
		blocks[tx.Height] = block.WithTxs{B: blocks[tx.Height].B, Txs: append(blocks[tx.Height].Txs, tx)}
	} else {
		blocks[tx.Height] = block.WithTxs{B: nil, Txs: []block.TxStatus{tx}}
	}

	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 18")
	// if we have filled the block with all necessary tx codes, send it to block stream and delete it from map
	// we need local copy because we might delete it earlier that block is getting send to stream
	localCopy := blocks[tx.Height]

	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 19")
	// we unlock here because full function could take some time to complete
	lock.Unlock()
	log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 20")
	if localCopy.Full() {
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 21")
		// send data to channel
		go func() { fullBlockStream <- localCopy }()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 22")
		// delete data which we sent to channel already
		lock.Lock()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 23")
		delete(blocks, tx.Height)
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 24")
		lock.Unlock()
		log.Println("pkg.cosmos_sdk.block.websocket.block_map.go - 25")
		return
	}
}
