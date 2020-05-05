package broker

import (
	watcher "github.com/mapofzones/cosmos-watcher/types"
)

type Broker struct {
	//todo
}

func New() (*Broker, error) {
	return &Broker{}, nil
}

func (broker *Broker) ExportBlock(block watcher.Block) error {
	//todo
	println("Block with height: ", block.Height, ". Done!")
	return nil
}
