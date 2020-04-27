package block

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/tendermint/tendermint/rpc/client/http"
)

// This function is used only to dump blocks to files in order to see if we are receiving them correctly
func TestBlockStream(t *testing.T) {
	client, err := http.New("tcp://localhost:26657", "/websocket")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Start()
	if err != nil {
		t.Fatal(err)
	}

	stream, err := GetBlockStream(context.Background(), client)
	if err != nil {
		t.Fatal(err)
	}

	for block := range stream {
		data, err := json.MarshalIndent(block, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(data))
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	}
}
