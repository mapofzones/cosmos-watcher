package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher/pkg"
	cosmos "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block"
	"github.com/mapofzones/cosmos-watcher/pkg/rabbitmq"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	// get height from which we should process blocks
	height, err := strconv.ParseInt(os.Getenv("height"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// initiate tendermint client for fetching blocks
	client, err := http.New(os.Getenv("rpc"), "/websocket")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Start()
	if err != nil {
		log.Fatal(err)
	}

	// initiate context for our app
	ctx, cancel := context.WithCancel(context.Background())

	// create block fetching pipeline
	blocks := cosmos.BlockStream(ctx, client, height)

	// initiate rabbitmq queue for watcher
	queue, err := rabbitmq.BlockQueue(ctx, os.Getenv("rabbitmq"), "blocks_v2")
	if err != nil {
		log.Fatal(err)
	}

	watcher := watcher.NewWatcher(ctx, blocks, queue)

	// run watcher
	err = watcher.WatchWithTimeout(ctx, time.Minute*10)
	cancel()
	log.Fatal(err)
}
