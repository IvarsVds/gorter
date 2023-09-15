package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type config struct {
	Directories map[string][]string `yaml:"directories"`
}

func readConfigFromFile() ([]byte, error) {
	var content []byte
	var err error

	// search config in 1) home folder, 2) global, 3) current folder
	// this seems rather ugly
	opsys := runtime.GOOS
	if opsys == "linux" {
		content, err = os.ReadFile(filepath.Join(os.Getenv("HOME"), "/.config/gorter.yaml"))
	}
	if opsys == "darwin" {
		content, err = os.ReadFile(filepath.Join(os.Getenv("HOME"), "/Library/Preferences/gorter.yaml"))
	}

	if err != nil {
		content, err = os.ReadFile("/etc/gorter.yaml")
	}

	if err != nil {
		content, err = os.ReadFile("./gorter.yaml")
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
