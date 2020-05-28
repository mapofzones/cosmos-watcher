package main

import (
	"context"
	"log"
	"os"

	watcher "github.com/mapofzones/cosmos-watcher/pkg"
)

func main() {
	watcher, err := watcher.NewWatcher(os.Getenv("rpc"), os.Getenv("rabbitmq"), os.Getenv("height"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	err = watcher.Watch(ctx)
	cancel()
	log.Fatal(err)
}
