package main

import (
	"encoding/csv"
	"github.com/RaphaelParment/metro/domain/models"
	"log"
	"os"
)

func loadInput(path string) (*models.Network, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %q; %v", path, err)
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'

	// Header
	_, err = r.Read()
	if err != nil {
		log.Printf("failed to read header; %v", err)
		return nil, err
	}

	return models.NewNetwork(r)
}

func run() error {
	network, err := loadInput("stations.csv")
	if err != nil {
		return err
	}
	for _, station := range network.Stations {
		log.Println(station.Name)
	}
	log.Println(len(network.Stations))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("main error; %v", err)
		os.Exit(1)
	}
}
