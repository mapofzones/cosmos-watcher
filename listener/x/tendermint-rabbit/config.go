package listener

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var configPath *string
var defualtPath = "/etc/event-listener.json"

func init() {
	configPath = flag.String("config-path", defualtPath, "path to a config file")
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
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		err = writeDefaultConfig()
		if err != nil {
			return nil, fmt.Errorf(
				"config file not found, could not write default default config to %s \n because %s",
				defualtPath, err.Error())
		}
		return nil, fmt.Errorf("could not parse config file, writing default one at %s", defualtPath)
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

func writeDefaultConfig() error {
	// create default config
	tmNode := url.URL{Scheme: "ws", Host: "localhost:26657", Path: "websocket"}
	rabbitNode := url.URL{Host: "localhost:5672"}
	defaultConfig := Config{NodeAddr: tmNode.String(),
		RabbitMQAddr: rabbitNode.String(),
		BatchSize:    1}

	bytes, err := json.MarshalIndent(defaultConfig, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(defualtPath, bytes, 0644)
}
