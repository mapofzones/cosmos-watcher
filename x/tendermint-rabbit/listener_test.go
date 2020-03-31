package listener

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestListenAndServe(t *testing.T) {
	logger := log.New(os.Stdout, "", log.Flags())
	l, err := NewListener(logger)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(l.ListenAndServe())
}

func TestNetworkName(t *testing.T) {
	url := url.URL{Scheme: "http", Host: "localhost:26657"}
	name, err := getNetworkName(url)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}
