package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/RaphaelParment/metro/domain/models"
	"log"
	"os"
)

func loadInput(path string) (*models.TransportNetwork, error) {
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

	return models.NewTransportNetwork(r)
}

func LoadStraightLine(number int) error {
	fileName := fmt.Sprintf("data/metro_%d.txt", number)
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("failed to open file %q; %v", fileName, err)
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stations []*models.MetroStation
	counter := 0
	var current *models.MetroStation

	for scanner.Scan() {
		name := scanner.Text()

		station := models.NewMetroStations(name, number)
		stations = append(stations, station)

		if counter == 0 {
			stations[counter].Prev[number] = nil
		} else {
			stations[counter].Prev[number] = current
		}

		current = station
		if counter > 0 {
			stations[counter-1].Next[number] = current
		}
		counter++
	}

	stations[counter-1].Next[number] = nil

	current = stations[0]
	for {
		if current == nil {
			break
		}
		log.Printf("name: %s\n", current.Name)
		current = current.Next[number]
	}

	return nil
}

func run() error {
	//transportNetwork, err := loadInput("stations.csv")
	//if err != nil {
	//	return err
	//}
	//for _, station := range transportNetwork.Stations {
	//	log.Println(station.Name)
	//}
	//log.Println(len(transportNetwork.Stations))

	return LoadStraightLine(1)


	//return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("main error; %v", err)
		os.Exit(1)
	}
}
