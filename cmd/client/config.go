package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type worker struct {
	Id            string `yaml:"id"`
	RemoteAddress string `yaml:"remote-address"`
}

type config struct {
	Workers []*worker `yaml:"workers"`
}

func getConfig() config {
	f, err := os.Open("client.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Error in reading client.yml ", err)
	}
	return cfg
}
