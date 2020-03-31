package main

import (
	"log"
	"os"

	watcher "github.com/attractor-spectrum/cosmos-watcher"
)

func main() {
	logger := log.New(os.Stdout, "Watcher", log.Ldate|log.Ltime)
	watcher, err := watcher.NewWatcher(watcher.TmRabbit, logger)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(watcher.Watch())
}
