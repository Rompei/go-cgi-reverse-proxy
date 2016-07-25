package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// Config is configuration of this program.
type Config struct {
	Server    string        `yaml:"server"`
	Port      uint          `yaml:"port"`
	WebRoot   string        `yaml:"webRoot"`
	ProxyRoot string        `yaml:"proxyRoot"`
	LogFile   string        `yaml:"logFile"`
	Path      []interface{} `yaml:"path"`
}

func loadConfig(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	if config.WebRoot, err = filepath.Abs(config.WebRoot); err != nil {
		return nil, err
	}
	if config.ProxyRoot, err = filepath.Abs(config.ProxyRoot); err != nil {
		return nil, err
	}
	return &config, nil
}
