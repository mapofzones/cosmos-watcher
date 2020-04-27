package watcher

import (
	"context"
	"fmt"
	"log"
	"testing"

	watcher "github.com/mapofzones/cosmos-watcher/types"
	"github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block"
)

func TestWatch(t *testing.T) {
	w, err := NewWatcher("tcp://localhost:26657", "amqp://guest@localhost:5672/")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(w.Watch(context.Background()))
}

func TestBlockStream(t *testing.T) {
	w, err := NewWatcher("tcp://localhost:26657", "amqp://guest@localhost:5672/")
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := block.GetBlockStream(context.Background(), w.client)
	if err != nil {
		log.Fatal(err)
	}

	for b := range blocks {
		fmt.Println(string(watcher.Normalize(b).JSON()))
	}
}
