package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Directories []Directory `yaml:"directories"`
}

type Directory struct {
	Name string   `yaml:"name"`
	Ext  []string `yaml:"ext"`
}

func readConfigFromFile() ([]byte, error) {
	var content []byte
	var err error

	// search config in 1) home folder, 2) global, 3) current folder
	home := os.Getenv("HOME")
	content, err = ioutil.ReadFile(filepath.Join(home, "/.config/gorter.yaml"))

	if err != nil {
		content, err = ioutil.ReadFile("/etc/gorter.yaml")
	}

	if err != nil {
		content, err = ioutil.ReadFile("./gorter.yaml")
	}

	if err != nil {
		err = errors.New("Couldn't load config file")
	}

	return content, err
}

func loadConfig() config {
	var config config

	content, err := readConfigFromFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Printf("Error while reading config file!\n %v\n", err)
		os.Exit(1)
	}

	return config
}
