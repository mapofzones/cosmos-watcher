package main

import (
	"flag"
	"github.com/mapofzones/cosmos-watcher/broker"
	"github.com/mapofzones/cosmos-watcher/client"
	"github.com/mapofzones/cosmos-watcher/processor"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	configsFlag  *string
	startHeight  int64
	latestHeight int64
	currentHeight int64
	lastExportedHeight int64
	workerCount  int16
	wg           sync.WaitGroup
)

func init() {
	configsFlag = flag.String("configs", "configs/", "path to a configs directory, where all files must be valid json config files")
	atomic.StoreInt64(&startHeight, 50)
	atomic.StoreInt64(&lastExportedHeight, atomic.LoadInt64(&startHeight)-1)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	workerCount = 1
}

func start() error {
	//os.Getenv()

	rpcNode := ""
	workers := make([]processor.Worker, workerCount, workerCount)
	cp, err := client.New(rpcNode)
	if err != nil {
		log.Error().Err(err).Msg("failed to start RPC client")
		return errors.Wrap(err, "failed to start RPC client")
	}
	br, err := broker.New()
	if err != nil {
		log.Error().Err(err).Msg("failed to start broker")
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
		log.Info().Int("number", i+1).Msg("starting worker...")
		go w.Start()
	}

	// listen for and trap any OS signal to gracefully shutdown and exit
	trapSignal()

	go enqueueMissingBlocks(exportQueue, cp)
	go startNewBlockListener(exportQueue, cp)

	// block main process (signal capture will call WaitGroup's Done)
	wg.Wait()
	return nil
}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the latestHeight up until the startHeight.
func enqueueMissingBlocks(exportQueue processor.Queue, cp client.ClientProxy) {
	var err error
	temporaryHeight, err := cp.LatestHeight()
	if err != nil {
		log.Error().Err(err).Msg( "failed to get lastest block from RPC client")
	}
	atomic.StoreInt64(&latestHeight, temporaryHeight)

	log.Info().Int64("latestHeight: ", atomic.LoadInt64(&latestHeight)).
		Int64(" startHeight: ", atomic.LoadInt64(&startHeight)).Msg("crawl started...")
	for {
		for atomic.StoreInt64(&currentHeight, atomic.LoadInt64(&startHeight));
		atomic.LoadInt64(&currentHeight) <= atomic.LoadInt64(&latestHeight); atomic.AddInt64(&currentHeight, 1) {
			if atomic.LoadInt64(&currentHeight) <= 0 {
				continue
			}
			log.Info().Int64("currentHeight", atomic.LoadInt64(&currentHeight)).Msg("crawler adds height to the queue")
			appendQueue(exportQueue, atomic.LoadInt64(&currentHeight))
		}
		time.Sleep(5 * time.Second)
		if atomic.LoadInt64(&lastExportedHeight) >= atomic.LoadInt64(&currentHeight)-1 {
			log.Info().Int64("currentHeight", atomic.LoadInt64(&currentHeight)).
				Int64("crawler processed height:", atomic.LoadInt64(&currentHeight)-1).Msg("crawler finished!")
			break
		}
		temporaryHeight, err := cp.LatestHeight()
		if err != nil {
			log.Error().Err(err).Msg("failed to get latest block from RPC client")
		}
		atomic.StoreInt64(&latestHeight, temporaryHeight)
		atomic.StoreInt64(&startHeight, atomic.LoadInt64(&currentHeight) + 1)
		log.Info().Int64("latestHeight", atomic.LoadInt64(&latestHeight)).Int64("startHeight", atomic.LoadInt64(&startHeight)).
			Int64(" currentHeight", atomic.LoadInt64(&currentHeight)).Msg("restart crawler!")
	}
}

// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(exportQueue processor.Queue, cp client.ClientProxy) {
	eventCh, cancel, err := cp.SubscribeNewBlocks("watcher-client")
	defer cancel()
	if err != nil {
		log.Error().Err(err).Msg("failed to subscribe to new blocks")
	}
	log.Info().Msg("listening for new block events...")

	for e := range eventCh {
		newBlock := e.Data.(tmtypes.EventDataNewBlock).Block
		height := newBlock.Header.Height
		if atomic.LoadInt64(&latestHeight) < atomic.LoadInt64(&currentHeight) {
			appendQueue(exportQueue, height)
		} else {
			atomic.StoreInt64(&latestHeight, height)
			log.Info().Int64("latestHeight", atomic.LoadInt64(&latestHeight)).Msg("has been increased")
		}
	}
}

func appendQueue(exportQueue processor.Queue, height int64) {
	if height == atomic.LoadInt64(&lastExportedHeight)+1 {
		log.Info().Int64("height", height).Msg("enqueueing missing block")
		exportQueue <- height
		atomic.AddInt64(&lastExportedHeight, 1)
	} else {
		log.Info().Int64("Wrong height:", height).Int64("must be 1 more than  lastExportedHeight:", atomic.LoadInt64(&lastExportedHeight))
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
		log.Info().Str("signal", sig.String()).Msg("caught signal; shutting down...")
		defer wg.Done()
	}()
}