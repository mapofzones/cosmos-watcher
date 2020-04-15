package tx

import (
	"encoding/json"
	"strings"
	"time"
	"unicode"

	"github.com/attractor-spectrum/cosmos-watcher/tx"
	"github.com/buger/jsonparser"
	fastJson "github.com/buger/jsonparser"
)

// Handle used by jsonparser
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
type Events map[string][]keyVal

// GetAttribute returns attribute of event, or empty string if event not found
func (e Events) GetAttribute(eventType, attrKey string) string {
	for _, a := range e[eventType] {
		if a.Key == attrKey {
			return a.Value
		}
	}
	return ""
}

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
}

// getEventsData parses binary json in order to find events that we need
func getEventsData(data []byte) ([]byte, error) {
	result, _, _, err := fastJson.Get(data, "result")
	if err != nil {
		return nil, ErrInvalidTx
	}
	events, _, _, err := fastJson.Get(result, "events")
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
		fastJson.ArrayEach(value, func(value []byte, dataType fastJson.ValueType, offset int, err error) {
			m[EventAndKey[0]] = append(m[EventAndKey[0]], keyVal{Key: EventAndKey[1], Value: string(value)})
		})
		return nil
	}
	fastJson.ObjectEach(data, populateMap)

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

// Normalize return transaction in it's non-blockchain specific form
func (t Tx) Normalize(txTime time.Time, network string, precision int) (stdTx tx.Tx) {
	switch t.Msg.Type {
	case "send":
		amount, denom := splitCoin(t.Msg.Events.GetAttribute("transfer", "amount"))
		stdTx.Sender = t.Msg.Events.GetAttribute("message", "sender")
		stdTx.Recipient = t.Msg.Events.GetAttribute("transfer", "recipient")
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Type = tx.Transfer
		stdTx.T = txTime
		stdTx.Network = network
		stdTx.Quantity = amount
		stdTx.Denom = denom
		stdTx.Precision = precision
		return
	case "create_validator":
		amount, denom := splitCoin(t.Msg.Events.GetAttribute("create_validator", "amount"))
		stdTx.Sender = t.Msg.Events.GetAttribute("message", "sender")
		stdTx.Quantity = amount
		stdTx.Denom = denom
		stdTx.Precision = precision
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.Stake
		return
	case "delegate":
		amount, denom := splitCoin(t.Msg.Events.GetAttribute("delegate", "amount"))
		stdTx.Sender = t.Msg.Events.GetAttribute("message", "sender")
		stdTx.Quantity = amount
		stdTx.Denom = denom
		stdTx.Precision = precision
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.Stake
		return
	case "begin_unbonding":
		amount, denom := splitCoin(t.Msg.Events.GetAttribute("unbond", "amount"))
		stdTx.Sender = t.Msg.Events.GetAttribute("message", "sender")
		stdTx.Quantity = amount
		stdTx.Denom = denom
		stdTx.Precision = precision
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.Unstake
		return
	case "transfer":
		stdTx = parseIbcSend(t)
		stdTx.Precision = precision
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.IbcSend
		return
	case "update_client", "ics20/transfer", "ics04/opaque":
		stdTx = parseIbcReceive(t)
		stdTx.Precision = precision
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.IbcRecieve
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		return
	default:
		stdTx.Sender = t.Msg.Events.GetAttribute("message", "sender")
		stdTx.Precision = precision
		stdTx.Hash = t.Msg.Events.GetAttribute("tx", "hash")
		stdTx.Network = network
		stdTx.T = txTime
		stdTx.Type = tx.Other
		data, _ := json.Marshal(t)
		stdTx.Data = data
		return
	}
}

func parseIbcSend(t Tx) tx.Tx {
	var packetData string
	for _, v := range t.Msg.Events["send_packet"] {
		if v.Key == "packet_data" {
			packetData = v.Value
			break
		}
	}
	var out tx.Tx
	// so we can parse raw string
	packetData = strings.Replace(packetData, "\\", "", -1)
	b := fastJson.StringToBytes(packetData)
	// get rid of amino encoding boilerplate
	b, _, _, err := fastJson.Get(b, "value")
	if err != nil {
		panic(err)
	}

	sender, err := fastJson.GetString(b, "sender")
	if err != nil {
		panic(err)
	}
	out.Sender = sender
	reciever, err := fastJson.GetString(b, "receiver")
	if err != nil {
		panic(err)
	}
	out.Recipient = reciever
	var amount, denom string
	// binary reprsentation of amount, which is array of amount-denom pairs
	amount, err = fastJson.GetString(b, "amount", "[0]", "amount")
	if err != nil {
		panic(err)
	}
	out.Quantity = amount
	denom, err = fastJson.GetString(b, "amount", "[0]", "denom")
	if err != nil {
		panic(err)
	}
	// since denom now has form a/b/c and we only need c
	denomSlice := strings.Split(denom, "/")
	denom = denomSlice[len(denomSlice)-1]
	out.Denom = denom
	return out
}

func parseIbcReceive(t Tx) tx.Tx {
	var sender string
	// this is how we decide who is the sender
	for _, v := range t.Msg.Events["message"] {
		if v.Key == "sender" {
			sender = v.Value
			break
		}
	}

	var recipient string
	// we have to get the recipient
	for _, v := range t.Msg.Events["transfer"] {
		if v.Key == "recipient" && v.Value != sender {
			recipient = v.Value
			break
		}
	}

	// get coin
	var coinStr string
	for _, v := range t.Msg.Events["transfer"] {
		if v.Key == "amount" {
			coinStr = v.Value
			break
		}
	}

	amount, denom := splitCoin(coinStr)

	return tx.Tx{Sender: sender, Recipient: recipient, Denom: denom, Quantity: amount}
}

// return amount, denom strings
func splitCoin(s string) (string, string) {
	for i, ch := range s {
		if !unicode.IsDigit(ch) {
			return s[:i], s[i:]
		}
	}
	panic("invalid amount string")
}
