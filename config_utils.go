package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config is configuration of this program.
type Config struct {
	Server  string        `yaml:"server"`
	Port    uint          `yaml:"port"`
	LogFile string        `yaml:"logFile"`
	Path    []interface{} `yaml:"path"`
}

func loadConfig(filename string) (config Config, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return
	}
	return
}
