package watcher

import (
	"fmt"
	"net/url"
	"testing"

	config "github.com/attractor-spectrum/cosmos-watcher/x/config"
)

func TestWatch(t *testing.T) {
	c, err := config.GetDefaultConfig()
	l, err := NewWatcher(c)
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
