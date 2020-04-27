package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher"
	config "github.com/mapofzones/cosmos-watcher/x/config"
)

var configsFlag *string

func init() {
	configsFlag = flag.String("configs", "configs/", "path to a configs directory, where all files must be valid json config files")
}

func configFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	res := make([]string, 0, len(files))
	for _, f := range files {
		if !f.IsDir() {
			file, _ := filepath.Abs(dir + f.Name())
			res = append(res, file)
		}
	}
	return res
}

func getConfigs(files []string) ([]*config.Config, error) {
	configs := make([]*config.Config, 0, len(files))
	for _, f := range files {
		config, err := config.GetConfig(f)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func main() {
	// just log raw data  without any prefixes
	log.SetFlags(0)

	flag.Parse()
	configs, err := getConfigs(configFiles(*configsFlag))
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, c := range configs {
		wg.Add(1)
		w, err := watcher.NewWatcher(watcher.TmRabbit, c)
		if err != nil {
			fmt.Println(err)
			// don't watch because we got networking error
			continue
		}
		go func() {
			defer wg.Done()
			t := time.Minute
			for {
				fmt.Println(w.Watch())
				time.Sleep(t)
				t += t
			}
		}()
	}
	wg.Wait()
}
