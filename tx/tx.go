package tx

import (
	"encoding/json"
	"time"
)

// Type tells us what tx we got
type Type string

const (
	Transfer   Type = "transfer"
	Stake      Type = "stake"
	Unstake    Type = "unstake"
	IbcSend    Type = "ibc-send"
	IbcRecieve Type = "ibc-recieve"
	Other      Type = "other"
)

// Tx represents transaction structure which is not blockchain specific
type Tx struct {
	T         time.Time `json:"time"`
	Hash      string    `json:"hash"`
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient,omitempty"`
	Quantity  string    `json:"quantity,omitempty"`
	Denom     string    `json:"denom,omitempty"`
	Network   string    `json:"network"`
	Type      Type      `json:"type"`
	Data      []byte    `json:"data,omitempty"`
	Precision int       `json:"precision,omitempty"`
}

// Marshal encodes our Tx with JSON encoding
func (t Tx) Marshal() []byte {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return data
}

// Txs is a wrapper around Tx slice providing neccessary utility functions
type Txs []Tx

// Marshal returns binary JSON representation of tx slice
func (t Txs) Marshal() []byte {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return data
}
