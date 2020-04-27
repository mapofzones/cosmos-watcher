package main

import (
	"context"
	"flag"
	"log"

	watcher "github.com/mapofzones/cosmos-watcher"
	types "github.com/mapofzones/cosmos-watcher/types"
)

var tendermintRPC, rabbitmqAddr *string

func init() {
	tendermintRPC = flag.String("tmRPC", "tcp://localhost:26657", "address of tendermint node to connect to")
	rabbitmqAddr = flag.String("rabbitMQ", "", "rabbitmq url to connect to")
}

func main() {
	// log raw data without time and date prefixes
	log.SetFlags(0)
	// parse flags from stdin
	flag.Parse()

	watcher, err := watcher.NewWatcher(types.TmRabbit, *tendermintRPC, *rabbitmqAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(watcher.Watch(context.Background()))
}
