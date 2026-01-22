package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Interval  int                 `yaml:"interval"`
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	URL      string `yaml:"url"`
	Protocol string `yaml:"protocol"`
}

func parseConfig() *Config {
	var config Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error while unmarshalling config file: %v", err)
	}

	fmt.Printf("Interval: %d\n", config.Interval)
	for key, ep := range config.Endpoints {
		fmt.Printf("Key: %s\n \tURL: %s\n \tProtocol: %s\n", key, ep.URL, ep.Protocol)
	}

	return &config
}
