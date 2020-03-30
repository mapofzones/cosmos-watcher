package listener

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

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
func (l *Listener) listen() <-chan Tx {
	// maybe want it buffered
	txs := make(chan Tx)

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
			l.logger.Println(err)
			close(txs)
			return
		}
		l.logger.Printf("established websocket connection with %s at %s\n", l.network, conn.RemoteAddr())
		// query that specifies what events we want from tendermint node
		err = conn.WriteMessage(websocket.TextMessage, txsQuery)
		if err != nil {
			l.logger.Println(err)
			close(txs)
			return
		}
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				l.logger.Println(err)
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

	return txs
}

// serve buffers and sends txs to rabbitmq, returns errors if something is wrong
func (l *Listener) serve(txs <-chan Tx) error {
	for tx := range txs {
		if tx.Valid {
			// l.logger.Printf("recieved valid %s tx", tx.Msg.Type)
			l.txs = append(l.txs, tx)
			if len(l.txs) == l.batchSize {
				go dummyImpl(l.txs)
				l.txs = make([]Tx, 0, l.batchSize)
			}
		}
	}
	return nil
}

// ListenAndServe implements listener interface
// Collects txs from tendermint websocket and sends them to rabbitMQ
func (l *Listener) ListenAndServe() error {
	txs := l.listen()
	return l.serve(txs)
}

func dummyImpl(txs []Tx) {
	data, _ := json.MarshalIndent(txs, "", "\t")
	fmt.Println(string(data))
}
