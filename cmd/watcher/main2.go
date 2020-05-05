package main

import (
	"flag"
	"github.com/mapofzones/cosmos-watcher/broker"
	"github.com/mapofzones/cosmos-watcher/client"
	"github.com/mapofzones/cosmos-watcher/processor"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	configsFlag *string
	lowerHeight int64
	biggerHeight int64
	workerCount int16
	wg          sync.WaitGroup
)

func init() {
	configsFlag = flag.String("configs", "configs/", "path to a configs directory, where all files must be valid json config files")
	lowerHeight = 50
	biggerHeight = 60
}

func start() error {

	workerCount := 1
	rpcNode := "<rpc node addr>"
	workers := make([]processor.Worker, workerCount, workerCount)
	cp, err := client.New(rpcNode)
	if err != nil {
		//todo
		println("failed to start RPC client")
		return errors.Wrap(err, "failed to start RPC client")
	}
	br, err := broker.New()
	if err != nil {
		return errors.Wrap(err, "failed to start broker")
	}
	exportQueue := processor.NewQueue(100)
	for i := range workers {
		workers[i] = processor.NewWorker(cp, br, exportQueue)
	}

	wg.Add(1)

	// Start each blocking worker in a go-routine where the worker consumes jobs
	// off of the export queue.
	for i, w := range workers {
		println("number", i+1, "starting worker...")
		go w.Start()
	}

	// listen for and trap any OS signal to gracefully shutdown and exit
	trapSignal()

	go enqueueMissingBlocks(exportQueue)

	// block main process (signal capture will call WaitGroup's Done)
	wg.Wait()

	return nil
}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the biggerHeight up until the lowerHeight.
func enqueueMissingBlocks(exportQueue processor.Queue) {
	//todo: loggin
	println("syncing missing blocks...")

	println("biggerHeight: ", biggerHeight, " lowerHeight: ", lowerHeight)
	for i := lowerHeight; i <= biggerHeight; i++ {
		println("index: ", i)
		if i == 1 {
			// skip the first block
			continue
		}
		//todo: logging

		println("height", i, "enqueueing missing block")
		exportQueue <- i
	}
}

func main() {
	// just log raw data  without any prefixes
	log.SetFlags(0)

	flag.Parse()

	start()
}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal() {
	var sigCh = make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		//todo: logging
		println("signal", sig.String(), "caught signal; shutting down...")
		defer wg.Done()
	}()
}