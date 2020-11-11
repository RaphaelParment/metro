package main

import (
	"bufio"
	"fmt"
	"github.com/RaphaelParment/metro/route_finder/domain/metro"
	"log"
	"os"
)

func loadStraightLine(network *metro.Network, number int) (*metro.Line, error) {
	fileName := fmt.Sprintf("data/metro_%d.txt", number)
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("failed to open file %q; %v", fileName, err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close %q; %v", fileName, err)
		}
	}()

	line := metro.NewLine(fmt.Sprint(number), number, nil)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stations []*metro.Station
	counter := 0
	var current *metro.Station

	for scanner.Scan() {
		var station *metro.Station
		name := scanner.Text()

		if s, ok := network.Stations[name]; !ok {
			station = metro.NewStation(name, number)
			network.Stations[name] = station
		} else {
			station = s
		}
		stations = append(stations, station)

		if counter > 0 {
			stations[counter].Neighbours = append(station.Neighbours, current)
			stations[counter].DistToNeighbours[current] = 1
		}

		current = station
		if counter > 0 {
			stations[counter-1].Neighbours = append(stations[counter-1].Neighbours, current)
			stations[counter-1].DistToNeighbours[current] = 1
		}
		counter++
		line.Stations = append(line.Stations, station)
	}

	return line, nil
}

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
