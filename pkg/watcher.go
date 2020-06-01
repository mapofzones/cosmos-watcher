package watcher

import (
	"context"
	"errors"
	"strconv"

	codec "github.com/mapofzones/cosmos-watcher/pkg/codec"
	cosmos "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block"
	"github.com/mapofzones/cosmos-watcher/pkg/rabbitmq"
	types "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/rpc/client/http"
)

// Watcher implements the watcher interface described at the project root
// this particular implementation is used to listen on tendermint websocket and
// send results to RabbitMQ
type Watcher struct {
	tendermintAddr string
	rabbitMQAddr   string
	client         *http.HTTP
	chainID        string
	startingHeight int64
}

// NewWatcher returns instanciated Watcher
func NewWatcher(tendermintRPCAddr, rabbitmqAddr string, startingHeight string) (*Watcher, error) {
	height, err := strconv.ParseInt(startingHeight, 10, 64)
	if err != nil {
		return nil, err
	}

	client, err := http.New(tendermintRPCAddr, "/websocket")
	if err != nil {
		return nil, err
	}
	err = client.Start()
	if err != nil {
		return nil, err
	}
	//figure out name of the blockchain to which we are connecting
	info, err := client.Status()
	if err != nil {
		return nil, err
	}

	return &Watcher{tendermintAddr: tendermintRPCAddr, rabbitMQAddr: rabbitmqAddr, chainID: info.NodeInfo.Network, client: client, startingHeight: height}, nil
}

// serve buffers and sends txs to rabbitmq, returns errors if something is wrong
// Accepts input and output channels, also an error channel if something goes wrong
func (w *Watcher) serve(ctx context.Context, blockInput <-chan types.Block, blockOutput chan<- types.Block) error {
	for {
		select {
		case block, ok := <-blockInput:
			if !ok {
				return errors.New("block channel is closed")
			}
			select {
			case blockOutput <- block:
			case <-ctx.Done():
			}
		case <-ctx.Done():
			err := w.client.Stop()
			w.client.Wait()
			return err
		}
	}
}

// Watch implements watcher interface
// Collects txs from tendermint websocket and sends them to rabbitMQ
func (w *Watcher) Watch(ctx context.Context) error {
	cdc := amino.NewCodec()
	codec.RegisterTypes(cdc)
	queue, err := rabbitmq.BlockQueue(ctx, w.rabbitMQAddr, "blocks_v2", cdc)
	if err != nil {
		return err
	}

	blocks := cosmos.BlockStream(ctx, w.client, w.startingHeight)

	return w.serve(ctx, blocks, queue)
}
