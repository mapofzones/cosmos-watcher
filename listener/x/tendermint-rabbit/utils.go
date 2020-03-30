package listener

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	json "github.com/buger/jsonparser"
)

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
