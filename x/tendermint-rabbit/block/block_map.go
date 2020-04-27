package block

import (
	"sync"
)

func init() {
	blocks = make(map[int64]WithTxs)
	lock = &sync.Mutex{}
}

// assume that each process is launched in separate binary
// so all records in map belong to one chain

// blocks is a global map object where tx data is stored
// should not be used directly, instead use helper methods in the package
var blocks map[int64]WithTxs
var lock *sync.Mutex

// pushBlock puts block in map and sends its over a full block channel for further parsing
func pushBlock(b Block) {
	// check if block has any transactions
	if len(b.Txs) == 0 {
		// send block for further processing
		go func() { fullBlockStream <- WithTxs{B: &b, Txs: []TxStatus{}} }()
		return
	}

	// our write to map critical section
	lock.Lock()
	// if WithTx structure is already present in the map
	if _, ok := blocks[b.Height]; ok {
		blocks[b.Height] = WithTxs{B: &b, Txs: blocks[b.Height].Txs}
	} else {
		blocks[b.Height] = WithTxs{B: &b, Txs: make([]TxStatus, len(b.Txs))}
	}

	// if we have filled the block with all necessary tx codes, send it to block stream and delete it from map
	// we need local copy because we might delete it earlier that block is getting send to stream
	localCopy := blocks[b.Height]

	// we unlock here because full function could take some time to complete
	lock.Unlock()
	if localCopy.Full() {
		// send data to channel
		go func() { fullBlockStream <- localCopy }()
		// delete data which we sent to channel already
		lock.Lock()
		delete(blocks, b.Height)
		lock.Unlock()
		return
	}
	return
}

// pushTx puts transaction in our map and, if block if full, sends it to block channel
func pushTx(tx TxStatus) {
	// our write to map critical section
	lock.Lock()
	// if WithTx structure is already present in the map
	if _, ok := blocks[tx.Height]; ok {
		blocks[tx.Height] = WithTxs{B: blocks[tx.Height].B, Txs: append(blocks[tx.Height].Txs, tx)}
	} else {
		blocks[tx.Height] = WithTxs{B: nil, Txs: []TxStatus{tx}}
	}

	// if we have filled the block with all necessary tx codes, send it to block stream and delete it from map
	// we need local copy because we might delete it earlier that block is getting send to stream
	localCopy := blocks[tx.Height]

	// we unlock here because full function could take some time to complete
	lock.Unlock()
	if localCopy.Full() {
		// send data to channel
		go func() { fullBlockStream <- localCopy }()
		// delete data which we sent to channel already
		lock.Lock()
		delete(blocks, tx.Height)
		lock.Unlock()
		return
	}
	return
}
