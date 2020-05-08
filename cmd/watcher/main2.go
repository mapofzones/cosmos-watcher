package main

import (
	"github.com/mapofzones/cosmos-watcher/broker"
	"github.com/mapofzones/cosmos-watcher/client"
	"github.com/mapofzones/cosmos-watcher/processor"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	startHeight  int64
	latestHeight int64
	currentHeight int64
	lastExportedHeight int64
	workerCount  int16
	wg           sync.WaitGroup
	rpcNode string
)

func getenvStr(key string) (string, error) {
	variable := os.Getenv(key)
	if variable == "" {
		return variable, errors.New("getenv: empty environment variable " + key)
	}
	return variable, nil
}

func getenvInt(key string) (int, error) {
	str, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	variable, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return variable, nil
}

func getenvIntSize(key string, size int) (int64, error) {
	str, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	variable, err := strconv.ParseInt(str, 10, size)
	if err != nil {
		return 0, err
	}
	return variable, nil
}

func getenvInt64(key string) (int64, error) {
	return getenvIntSize(key, 64)
}

func getenvInt16(key string) (int16, error) {
	variable, err := getenvIntSize(key, 16)
	return int16(variable), err
}

func getenvBool(key string) (bool, error) {
	str, err := getenvStr(key)
	if err != nil {
		return false, err
	}
	variable, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return variable, nil
}

func init() {
	//os.Setenv("rpc", "")
	//os.Setenv("startHeight", "1")
	//os.Setenv("workerCount", "1")

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	height, err := getenvInt64("startHeight")
	if err != nil {
		log.Info().Msg("The value of startHeight is not set, the default value of 1 is used")
		atomic.StoreInt64(&startHeight, 1)
	} else {
		atomic.StoreInt64(&startHeight, height)
	}
	atomic.StoreInt64(&lastExportedHeight, atomic.LoadInt64(&startHeight)-1)

	workerCount, err = getenvInt16("workerCount")
	if err != nil {
		log.Info().Msg("The value of workerCount is not set, the default value of 1 is used")
		workerCount = 1
	}
	rpcNode, err = getenvStr("rpc")
	if err != nil {
		log.Error().Err(err).Msg("failed to parse RPC env variable")
	}
}

func start() error {
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