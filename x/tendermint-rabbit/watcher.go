package watcher

import (
	"bytes"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/attractor-spectrum/cosmos-watcher/tx"
	config "github.com/attractor-spectrum/cosmos-watcher/x/config"
	rabbitmq "github.com/attractor-spectrum/cosmos-watcher/x/tendermint-rabbit/RabbitMQ"
	txparser "github.com/attractor-spectrum/cosmos-watcher/x/tendermint-rabbit/tx"
	"github.com/gorilla/websocket"
)

// txsQuery specifying that we want all transactions
var txsQuery = []byte("{\"jsonrpc\":\"2.0\",\"id\":0,\"method\":\"subscribe\",\"params\":{\"query\":\"tm.event = 'Tx'\"}}\n")

// Tx is just an alias for txparser.Tx type
type Tx = tx.Tx

// Txs is an alias for our convience
type Txs = tx.Txs

// Watcher implenets the watcher interface described at the project root
// this particular implementation is used to listen on tendermint websocket and
// send resutls to RabbitMQ
type Watcher struct {
	tendermintAddr url.URL
	rabbitMQAddr   url.URL
	// how many txs we accumulate before sending them for further processing
	batchSize int
	network   string
	txs       Txs
	precision int
	config    *config.Config
}

// NewWatcher returns instanciated Watcher
func NewWatcher(config *config.Config) (*Watcher, error) {
	//we checked if urls are valid in GetConfig already
	nodeURL, _ := url.Parse(config.NodeAddr)
	rabbitURL, _ := url.Parse(config.RabbitMQAddr)

	//figure out name of the blockchain to which we are connecting
	name, err := getNetworkName(*nodeURL)
	if err != nil {
		return nil, err
	}
	// create tx slice with our config capacity
	txs := make([]Tx, 0, config.BatchSize)
	return &Watcher{tendermintAddr: *nodeURL, rabbitMQAddr: *rabbitURL,
		network: name, txs: txs, batchSize: config.BatchSize, precision: config.Precision, config: config}, nil
}

// listen creates goroutine which reads txs from a websocket and pushes them to Tx channel
// we also need err channel to know if something happened to our websocket connection
func (l *Watcher) listen() (<-chan Tx, <-chan error) {
	// maybe want it buffered
	txs := make(chan Tx)
	errCh := make(chan error)

	go func() {
		conn, _, err := websocket.DefaultDialer.Dial(l.tendermintAddr.String(), nil)
		// close connection on exit
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()
		// we could not establish the connection
		if err != nil {
			errCh <- err
			close(errCh)
			close(txs)
			return
		}
		fmt.Printf("established websocket connection with %s at %s\n", l.network, conn.RemoteAddr())
		// query that specifies what events we want from tendermint node
		err = conn.WriteMessage(websocket.TextMessage, txsQuery)
		if err != nil {
			errCh <- err
			close(txs)
			close(errCh)
			return
		}
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				errCh <- err
				close(errCh)
				close(txs)
				return
			}
			// make it a tx and send it to querier
			go func() {
				t := time.Now().UTC()
				// we got normal json rpc greeting, do nothing
				if bytes.Equal(data, rpcGreeting) {
					return
				}
				// send it for debbuging
				DebugSend(l.rabbitMQAddr, DebugData{Data: data, Chain: l.network, Time: time.Now()})
				tmTx, err := txparser.ParseTx(data)
				if !tmTx.Valid {
					fmt.Printf("recieved invalid tx: %v", tmTx)
					return
				}
				if err != nil {
					fmt.Printf("expected tendermint tx, got: %s\n%v", string(data), err)
					return
				}

				txs <- tmTx.Normalize(t, l.network, l.precision)
			}()
		}

	}()

	return txs, errCh
}

// serve buffers and sends txs to rabbitmq, returns errors if something is wrong
// Accepts input and output channels, also an error channel if something goes wrong
func (l *Watcher) serve(txsIn <-chan Tx, txsOut chan<- Txs, errors <-chan error) error {
	for {
		select {
		case tx := <-txsIn:
			fmt.Printf("recieved valid cosmos-sdk %s tx\n", tx.Type)
			l.txs = append(l.txs, tx)
			if len(l.txs) == l.batchSize {
				select {
				case txsOut <- l.txs:
					l.txs = make([]Tx, 0, l.batchSize)
				case err := <-errors:
					return err
				}
			}
		case err := <-errors:
			return err
		}
	}
}

// Watch implements watcher interface
// Collects txs from tendermint websocket and sends them to rabbitMQ
func (l *Watcher) Watch() error {
	txsIn, errIn := l.listen()
	txsOut, errOut, err := rabbitmq.TxQueue(l.rabbitMQAddr)
	if err != nil {
		return err
	}
	errors := fanInErrors(errIn, errOut)
	return l.serve(txsIn, txsOut, errors)
}

func fanInErrors(errors ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	multiplexedErrors := make(chan error)

	multiplex := func(e <-chan error) {
		defer wg.Done()
		for err := range e {
			multiplexedErrors <- err
		}
	}

	wg.Add(len(errors))
	for _, e := range errors {
		go multiplex(e)
	}

	go func() {
		wg.Wait()
		close(multiplexedErrors)
	}()

	return multiplexedErrors
}
