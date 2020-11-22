package main

import (
	"bufio"
	"fmt"
	"github.com/RaphaelParment/metro/route_finder/domain/metro"
	"log"
	"os"
)



func run() error {
	network := metro.NewNetwork()
	line1, err := loadStraightLine(network, 1)
	if err != nil {
		return err
	}
	line2, err := loadStraightLine(network, 2)
	if err != nil {
		return err
	}

	lines := []*metro.Line{line1, line2}
	network.Lines = lines

	sps := metro.NewShortestPathSolver(network)
	source, _ := network.GetStation("La Defense - Grande Arche")
	dest, _ := network.GetStation("Belleville")
	sps.Solve(source, dest)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("main error; %v", err)
		os.Exit(1)
	}
}
