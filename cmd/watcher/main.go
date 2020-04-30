package main

import (
	"context"
	"flag"
	"log"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher"
	types "github.com/mapofzones/cosmos-watcher/types"
)

var tendermintRPC, rabbitmqAddr, zoneName *string

func init() {
	tendermintRPC = flag.String("tmRPC", "tcp://localhost:26657", "address of tendermint node to connect to")
	rabbitmqAddr = flag.String("rabbitMQ", "", "rabbitmq url to connect to")
	zoneName = flag.String("zone", "", "zone name to be displayed by logger")
}

func main() {
	// log raw data without time and date prefixes
	log.SetFlags(0)
	log.SetPrefix(*zoneName)
	// parse flags from stdin
	flag.Parse()
	run()
}

var t = time.Minute

func run() {
	for {
		then := time.Now()
		w, err := watcher.NewWatcher(types.TmRabbit, *tendermintRPC, *rabbitmqAddr)
		if err != nil {
			log.Println(err)
			goto Sleepy
		} else {
			ctx, cancel := context.WithCancel(context.Background())
			log.Println(w.Watch(ctx))
			cancel()
		}
	Sleepy:
		if time.Since(then).Minutes() > 10 {
			t = time.Minute
		} else {
			t = getDelay()
		}
		log.Println("could not connect to ", *tendermintRPC)
		time.Sleep(t)
		continue
	}

}

func getDelay() time.Duration {
	switch t {
	case time.Minute:
		return time.Minute * 10
	case time.Minute * 10:
		return time.Minute * 30
	case time.Minute * 30:
		return time.Hour
	case time.Hour:
		return t
	}
	return t
}
