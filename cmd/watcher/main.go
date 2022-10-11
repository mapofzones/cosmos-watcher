package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher/pkg"
	cosmos "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block"
	"github.com/mapofzones/cosmos-watcher/pkg/rabbitmq"
	"github.com/okex/exchain/libs/tendermint/rpc/client/http"
)

func main() {
	startWithBlockchainHeight := "14572137" //os.Getenv("height")
	//fullNodeJsonRpcAddress := "http://35.73.150.207:26657"                //os.Getenv("rpc")
	fullNodeJsonRpcAddress := "https://exchaintmrpc.okex.org"             //os.Getenv("rpc")
	messageBrokerConnectionString := "amqp://guest:guest@localhost:5672/" //os.Getenv("rabbitmq")
	rabbitmqQueueName := "watcher"                                        //os.Getenv("queue")
	blockchainNetworkId := "exchain-66"                                   //os.Getenv("chain_id")

	// get height from which we should process blocks
	height, err := strconv.ParseInt(startWithBlockchainHeight, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// initiate tendermint client for fetching blocks

	client, err := http.New(fullNodeJsonRpcAddress, "/websocket")
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
	info, err := client.Status()
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
