package watcher

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/url"
)

var configPath *string

const examplePath = "/abs/path/to/config.json"

func init() {
	configPath = flag.String("config", examplePath, "path to a config file, example config can be found in the project source code")
}

// Config has addresses to connect to and stores batch query size
// default config is stored in /etc
type Config struct {
	NodeAddr     string
	RabbitMQAddr string
	BatchSize    int
}

// GetConfig returns config or error, signifying invalid config
func GetConfig() (*Config, error) {
	if *configPath == examplePath {
		return nil, errors.New("no config file specified")
	}
	bytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}
	C := new(Config)
	err = json.Unmarshal(bytes, C)
	if err != nil {
		return nil, err
	}

	// some basic error checking
	if _, err := url.Parse(C.NodeAddr); err != nil {
		return nil, err
	}
	if _, err := url.Parse(C.RabbitMQAddr); err != nil {
		return nil, err
	}
	if C.BatchSize < 0 {
		return nil, errors.New("BatchSize must be positive")
	}
	return C, nil
}
