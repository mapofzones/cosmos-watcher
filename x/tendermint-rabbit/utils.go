package watcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	json "github.com/buger/jsonparser"
)

var rpcGreetingStr string = `{
  "jsonrpc": "2.0",
  "id": 0,
  "result": {}
}`

// this is only used so we don't log greeting message as an error
var rpcGreeting []byte = []byte(rpcGreetingStr)

// return tendermint network name
func getNetworkName(u url.URL) (string, error) {
	u.Path = "status"
	u.Scheme = "http"
	res, err := http.DefaultClient.Get(u.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid response code for query %s", u.String())
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	network, _, _, err := json.Get(bytes, "result", "node_info", "network")
	if err != nil {
		return "", err
	}
	return string(network), nil
}
