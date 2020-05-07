package main

import (
	"flag"
	"github.com/mapofzones/cosmos-watcher/broker"
	"github.com/mapofzones/cosmos-watcher/client"
	"github.com/mapofzones/cosmos-watcher/processor"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	configsFlag  *string
	startHeight  int64
	latestHeight int64
	currentHeight int64
	workerCount  int16
	lastExportedHeight int64
	wg           sync.WaitGroup
)

func init() {
	configsFlag = flag.String("configs", "configs/", "path to a configs directory, where all files must be valid json config files")
	startHeight = 50
	lastExportedHeight = startHeight-1
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

	go enqueueMissingBlocks(exportQueue, cp)
	go startNewBlockListener(exportQueue, cp)
	//time.Sleep(10 * time.Second)
	// block main process (signal capture will call WaitGroup's Done)
	wg.Wait()
	return nil
}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the latestHeight up until the startHeight.
func enqueueMissingBlocks(exportQueue processor.Queue, cp client.ClientProxy) {
	latestHeight, err := cp.LatestHeight()
	if err != nil {
		log.Fatal().Err(errors.Wrap(err, "failed to get lastest block from RPC client"))
	}
	//todo: loggin
	println("syncing missing blocks...")
	println("latestHeight: ", latestHeight, " startHeight: ", startHeight)
	for {
		for currentHeight := startHeight; currentHeight <= latestHeight; currentHeight++ {
			println("index: ", currentHeight)
			if currentHeight == 1 {
				// skip the first block
				continue
			}
			//todo: logging
			appendQueue(exportQueue, currentHeight)
		}
		if lastExportedHeight >= currentHeight {
			break
		}
	}
}

// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(exportQueue processor.Queue, cp client.ClientProxy) {
	eventCh, cancel, err := cp.SubscribeNewBlocks("watcher-client")
	defer cancel()

	if err != nil {
		log.Fatal().Err(errors.Wrap(err, "failed to subscribe to new blocks"))
	}

	log.Info().Msg("listening for new block events...")

	for e := range eventCh {
		newBlock := e.Data.(tmtypes.EventDataNewBlock).Block
		height := newBlock.Header.Height

		if height > currentHeight {
			latestHeight = height
			log.Info().Int64("latestHeight", latestHeight).Msg(" has been increased")
		} else {
			appendQueue(exportQueue, currentHeight)
		}
	}
}

func appendQueue(exportQueue processor.Queue, height int64) {
	if height == lastExportedHeight+1 {
		println("height", height, "enqueueing missing block")
		exportQueue <- height
		lastExportedHeight++
	} else {
		println("Wrong height:", height, " must be 1 more than  lastExportedHeight:", lastExportedHeight)
		os.Exit(0)
	}
}

func main() {
	// just log raw data  without any prefixes
	//log.SetFlags(0)

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