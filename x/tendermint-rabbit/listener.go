package listener

import (
	"flag"
	"log"
	"net/url"
	"sync"
	"time"

	rabbitmq "github.com/attractor-spectrum/cosmos-watcher/listener/x/tendermint-rabbit/RabbitMQ"
	txparser "github.com/attractor-spectrum/cosmos-watcher/listener/x/tendermint-rabbit/tx"
	"github.com/gorilla/websocket"
)

// txsQuery specifying that we want all transactions
var txsQuery = []byte("{\"jsonrpc\":\"2.0\",\"id\":0,\"method\":\"subscribe\",\"params\":{\"query\":\"tm.event = 'Tx'\"}}\n")

// Tx is just an alias for txparser.Tx type
type Tx = txparser.Tx

// Listener implenets the listener interface described at the project root
// this particular implementation is used to listen on tendermint websocket and
// send resutls to RabbitMQ
type Listener struct {
	tendermintAddr url.URL
	rabbitMQAddr   url.URL
	logger         *log.Logger
	// how many txs we accumulate before sending them for further processing
	batchSize int
	network   string
	txs       []Tx
}

// NewListener returns instanciated Listener
func NewListener(l *log.Logger) (*Listener, error) {
	// figure out config
	// error out if there is no config
	flag.Parse()
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

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
	return &Listener{tendermintAddr: *nodeURL, rabbitMQAddr: *rabbitURL,
		logger: l, network: name, txs: txs, batchSize: config.BatchSize}, nil
}

// listen creates goroutine which reads txs from a websocket and pushes them to Tx channel
// we also need err channel to know if something happened to our websocket connection
func (l *Listener) listen() (<-chan Tx, <-chan error) {
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
		l.logger.Printf("established websocket connection with %s at %s\n", l.network, conn.RemoteAddr())
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
				tx, err := txparser.ParseTx(data)
				if err != nil {
					l.logger.Printf("expected tendermint tx, got: %s\n%v", string(data), err)
					return
				}
				tx.T = t
				txs <- tx
			}()
			// may want to implement heartbeat pattern
		}

	}()

	return txs, errCh
}

// serve buffers and sends txs to rabbitmq, returns errors if something is wrong
// Accepts input and output channels, also an error channel if something goes wrong
func (l *Listener) serve(txsIn <-chan Tx, txsOut chan<- []Tx, errors <-chan error) error {
	for {
		select {
		case tx := <-txsIn:
			if tx.Valid {
				l.logger.Printf("recieved valid %s tx", tx.Msg.Type)
				l.txs = append(l.txs, tx)
				if len(l.txs) == l.batchSize {
					select {
					case txsOut <- l.txs:
						l.txs = make([]Tx, 0, l.batchSize)
					case err := <-errors:
						return err
					}
				}
			}
		case err := <-errors:
			return err
		}
	}
}

// ListenAndServe implements listener interface
// Collects txs from tendermint websocket and sends them to rabbitMQ
func (l *Listener) ListenAndServe() error {
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
