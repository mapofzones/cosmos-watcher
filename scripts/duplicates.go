package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	type config struct {
		X string `json:"NodeAddr"`
	}

	files, err := ioutil.ReadDir("./configs")
	if err != nil {
		log.Fatal(err)
	}

	addresses := []string{}
	for _, f := range files {
		data, err := ioutil.ReadFile("./configs/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		x := config{}
		json.Unmarshal(data, &x)
		u, err := url.Parse(x.X)
		if err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, "tcp"+"://"+u.Host)
	}

	x := map[string][]string{}

	for _, addr := range addresses {
		client, err := http.New(addr, "/websocket")
		if err != nil {
			log.Fatal(err)
		}
		err = client.Start()
		if err != nil {
			log.Println("could not connect to: ", addr)

			continue
		}
		fmt.Println("trying to connect to ", addr)
		info, err := client.Status()
		if err != nil {
			log.Fatal(err)
		}
		x[info.NodeInfo.Network] = append(x[info.NodeInfo.Network], addr)

	}

	for key, value := range x {
		if len(value) > 1 {
			fmt.Println(key, " : ", value)
		}
	}
}

// atomicbombers  :  [tcp://goz.kalpatech.co:26657 tcp://ibc.izo.ro:26657]
