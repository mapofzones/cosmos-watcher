package main

import (
	"log"

	watcher "github.com/attractor-spectrum/cosmos-watcher"
	config "github.com/attractor-spectrum/cosmos-watcher/x/config"
)

func main() {
	config, err := config.GetDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := watcher.NewWatcher(watcher.TmRabbit, config)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(watcher.Watch())
}
