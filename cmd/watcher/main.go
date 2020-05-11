package main

import (
	"context"
	"log"
	"os"

	watcher "github.com/mapofzones/cosmos-watcher/pkg"
)

func main() {
	watcher, err := watcher.NewWatcher(os.Getenv("tmRPC"), os.Getenv("rabbitmq"), os.Getenv("StartHeight"))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(watcher.Watch(context.Background()))
}
