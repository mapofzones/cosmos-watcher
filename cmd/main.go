package main

import (
	"log"

	watcher "github.com/attractor-spectrum/cosmos-watcher"
)

func main() {
	watcher, err := watcher.NewWatcher(watcher.TmRabbit)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(watcher.Watch())
}
