package watcher

import (
	"context"
	"fmt"
	"net/url"
	"testing"
)

func TestWatch(t *testing.T) {
	w, err := NewWatcher("tcp://localhost:26657", "amqp://guest@localhost:5672/")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(w.Watch(context.Background()))
}

func TestNetworkName(t *testing.T) {
	url := url.URL{Scheme: "http", Host: "localhost:26657"}
	name, err := getNetworkName(url)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}
