package main

import (
	"context"
	"github.com/RaphaelParment/metro/loader/domain"
	"github.com/RaphaelParment/metro/loader/infrastructure"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("main: %v", err)
		os.Exit(1)
	}
}

func run() error {
	n := new(domain.Network)
	n.Stations = make(map[string]*domain.Station)

	l := new(infrastructure.FileLoader)
	regularLines := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 14}
	for _, line := range regularLines {
		if err := l.Load(context.TODO(), n, line); err != nil {
			log.Printf("fialed to load line %d; %v", line, err)
			return err
		}
	}

	for name, station := range n.Stations {
		log.Printf("name: %s -> lines %v", name, station.Lines)
		for i, neighbour := range station.Neighbours {
			log.Printf("\tn%d -> %s (%v)", i, neighbour.Name, neighbour.Lines)
		}
	}

	return nil
}
