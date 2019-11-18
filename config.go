package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Directories []directory
}

type directory struct {
	Name string   `json:"name"`
	Ext  []string `json:"ext"`
}

func readConfig() ([]byte, error) {
	var content []byte
	var err error

	// check for user config
	home := os.Getenv("HOME")
	content, err = ioutil.ReadFile(filepath.Join(home, "/.config/gorter_config.json"))

	if err != nil {
		// check for global config
		content, err = ioutil.ReadFile("/etc/gorter_config.json")
	}

	if err != nil {
		// check for config file in current folder
		content, err = ioutil.ReadFile("./gorter_config.json")
	}

	if err != nil {
		err = errors.New("Couldn't load config file")
	}

	return content, err
}

func loadConfig() config {
	var config config

	content, err := readConfig()

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(content, &config)
	return config
}
