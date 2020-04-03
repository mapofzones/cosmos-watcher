package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	watcher "github.com/attractor-spectrum/cosmos-watcher"
	config "github.com/attractor-spectrum/cosmos-watcher/x/config"
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
		}
		go func() {
			defer wg.Done()
			fmt.Println(w.Watch())
		}()
	}

	wg.Wait()
}
