package processor

import (
	"github.com/mapofzones/cosmos-watcher/broker"
	"github.com/mapofzones/cosmos-watcher/client"
	"github.com/rs/zerolog/log"
)

type (
	// Queue is a simple type alias for a (buffered) channel of block heights.
	Queue chan int64

	// Worker defines a job consumer that is responsible for getting and
	// aggregating block and associated data and exporting it to a database.
	Worker struct {
		cp    *client.ClientProxy
		broker *broker.Broker
		queue Queue
	}
)

func NewQueue(size int) Queue {
	return make(chan int64, size)
}

func NewWorker(cp *client.ClientProxy, broker *broker.Broker, queue Queue) Worker {
	return Worker{cp, broker,queue}
}

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {
	for i := range w.queue {
		log.Info().Int64("height", i).Msg("processing block")

		if err := w.process(i); err != nil {
			// re-enqueue any failed job
			// TODO: Implement exponential backoff or max retries for a block height.
			go func() {
				log.Info().Int64("height", i).Msg("re-enqueueing failed block")
				w.queue <- i
			}()
		}
	}
}

// process defines the job consumer workflow. It will fetch a block for a given
// height and associated metadata and export it to a broker. It returns an
// error if any export process fails.
func (w Worker) process(height int64) error {
	block, err := w.cp.GetBlockByHeight(height)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to get block")
		return err
	}

	txs, err := w.cp.GetTxsByBlock(block)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to get txs")
		return err
	}

	watcherBlock, err := w.cp.CreateWatcherBlock(block, txs)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to create watcher block")
		return err
	}

	err = w.broker.ExportBlock(watcherBlock)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to export block")
		return err
	}
	return nil
}
