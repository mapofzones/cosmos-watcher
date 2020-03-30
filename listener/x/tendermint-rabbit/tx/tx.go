package tx

import (
	"strings"
	"time"

	"github.com/buger/jsonparser"
	json "github.com/buger/jsonparser"
)

// Handle used jsonparser
type jsonFunc = func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error

// makes the whole Event thing less of a mess
type keyVal struct {
	Key   string
	Value string
}

// MessageType returns message type, empty strings means that tx has failed
func MessageType(attributes []keyVal) string {
	for _, v := range attributes {
		if v.Key == "action" {
			return v.Value
		}
	}
	return ""
}

// Events are meant to represent cosmos-sdk tx events structure
type Events = map[string][]keyVal

// Message represents cosmos-sdk message
type Message struct {
	Events Events `json:"events"`
	Type   string `json:"type"`
}

// Tx represents tendermint transaction
// Valid means that it somehow changes machine's state
type Tx struct {
	Valid bool    `json:"-"`
	Msg   Message `json:"tx"`
	// T stands for time when we received tx from websocket
	T time.Time `json:"time"`
}

// getEventsData parses binary json in order to find events that we need
func getEventsData(data []byte) ([]byte, error) {
	result, _, _, err := json.Get(data, "result")
	if err != nil {
		return nil, ErrInvalidTx
	}
	events, _, _, err := json.Get(result, "events")
	if err != nil {
		return nil, ErrInvalidTx
	}
	return events, nil
}

// creates events structure by parsing json bytes
// assumes data to be valid
func createEventsMap(data []byte) Events {
	m := make(Events)

	// parse each event and fill our map
	// lambda inside lambda stuff, don't think about it much
	populateMap := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		EventAndKey := strings.Split(string(key), ".")
		json.ArrayEach(value, func(value []byte, dataType json.ValueType, offset int, err error) {
			m[EventAndKey[0]] = append(m[EventAndKey[0]], keyVal{Key: EventAndKey[1], Value: string(value)})
		})
		return nil
	}
	json.ObjectEach(data, populateMap)

	return m
}

// ParseTx returns tx object and error
//  tx should be checked for if it is successdfull
func ParseTx(data []byte) (Tx, error) {
	events, err := getEventsData(data)
	if err != nil {
		return Tx{Valid: false}, err
	}

	m := createEventsMap(events)
	txType := MessageType(m["message"])

	// tx must have message.action, if it does not, that means that tx has failed
	if txType == "" {
		return Tx{Valid: false}, nil
	}
	return Tx{Valid: true, Msg: Message{Events: m, Type: txType}}, nil
}
