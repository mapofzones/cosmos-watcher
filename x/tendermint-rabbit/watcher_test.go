package watcher

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestWatch(t *testing.T) {
	logger := log.New(os.Stdout, "", log.Flags())
	l, err := NewWatcher(logger)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(l.Watch())
}

func TestNetworkName(t *testing.T) {
	url := url.URL{Scheme: "http", Host: "localhost:26657"}
	name, err := getNetworkName(url)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}
