package watcher

import (
	"fmt"
	"net/url"
	"testing"
)

func TestWatch(t *testing.T) {
	l, err := NewWatcher()
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
