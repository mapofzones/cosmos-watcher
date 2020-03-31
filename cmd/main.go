package main

import (
	"log"
	"os"

	"github.com/attractor-spectrum/cosmos-watcher/listener"
)

func main() {
	logger := log.New(os.Stdout, "Listener", log.Ldate|log.Ltime)
	listener, err := listener.NewListener(listener.TmRabbit, logger)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(listener.ListenAndServe())
}
