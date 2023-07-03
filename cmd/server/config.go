package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type config struct {
	Server struct {
		ListenAddress string `yaml:"listen"`
	} `yaml:"server"`
}

func getConfig() config {
	f, err := os.Open("config.yml")
	if err != nil {

	}
	defer f.Close()

	var cfg config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Error in reading config.yml ", err)
	}
	return cfg
}
