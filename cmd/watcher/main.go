package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cometbft/cometbft/rpc/client/http"
	watcher "github.com/mapofzones/cosmos-watcher/pkg"
	cosmos "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block"
	"github.com/mapofzones/cosmos-watcher/pkg/rabbitmq"
)

func main() {
	startWithBlockchainHeight := os.Getenv("height")
	fullNodeJsonRpcAddress := os.Getenv("rpc")
	messageBrokerConnectionString := os.Getenv("rabbitmq")
	rabbitmqQueueName := os.Getenv("queue")
	blockchainNetworkId := os.Getenv("chain_id")

	// get height from which we should process blocks
	height, err := strconv.ParseInt(startWithBlockchainHeight, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// initiate tendermint client for fetching blocks
	client, err := http.New(fullNodeJsonRpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Start()
	if err != nil {
		log.Println(err)
	}

	// initiate context for our app
	ctx, cancel := context.WithCancel(context.Background())

	// validate fullnode address
	info, err := client.Status(ctx)
	if err != nil {
		log.Fatal(err)
	} else if !strings.EqualFold(info.NodeInfo.Network, blockchainNetworkId) {
		log.Fatalf("Required chain_id(%s) is not equal to network_id(%s) in current blockchain fullnode", blockchainNetworkId, info.NodeInfo.Network)
	}

	// create block fetching pipeline
	blocks := cosmos.BlockStream(ctx, client, height)

	// initiate rabbitmq queue for watcher
	queue, err := rabbitmq.BlockQueue(ctx, messageBrokerConnectionString, rabbitmqQueueName)
	if err != nil {
		log.Fatal(err)
	}

	watcher := watcher.NewWatcher(ctx, blocks, queue)

	// run watcher
	err = watcher.WatchWithTimeout(ctx, time.Minute*10)
	cancel()
	log.Fatal(err)
}
