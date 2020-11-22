package infrastructure

import (
	"bufio"
	"context"
	"fmt"
	"github.com/RaphaelParment/metro/loader/domain"
	"log"
	"os"
)

type FileLoader struct {
}

func (l *FileLoader) Load(ctx context.Context, n *domain.Network, number string) error {
	switch number {
	case "7":
		err := l.loadLine7(n)
		if err != nil {
			return err
		}
		return nil
	case "13":
		err := l.loadLine13(n)
		if err != nil {
			return err
		}
		return nil
	default:
		records, err := l.readInput(number)
		if err != nil {
			return err
		}
		return l.loadStraightLine(n, records, number)
	}
}

func (l *FileLoader) readInput(number string) ([]string, error) {
	path := fmt.Sprintf("../data/metro%s.txt", number)
	f, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %q; %v", path, err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close %q; %v", path, err)
		}
	}()

	var records []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (l *FileLoader) loadStraightLine(n *domain.Network, stations []string, number string) error {
	counter := 0
	var previous *domain.Station
	for _, record := range stations {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, number)
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, number)
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}
	return nil
}

func (l *FileLoader) loadLine7(n *domain.Network) error {
	// Read metro7_1
	recordsLine7, err := l.readInput("7_1")
	if err != nil {
		return err
	}
	counter := 0
	var previous *domain.Station
	for _, record := range recordsLine7 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "7_1")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "7_1")
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	fork := previous

	// Read metro7_2
	recordsLine71, err := l.readInput("7_2")
	if err != nil {
		return err
	}
	counter = 0
	for _, record := range recordsLine71 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "7_2")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "7_2")
		}

		if counter == 0 {
			station.AddNeighbour(fork)
			fork.AddNeighbour(station)
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	// Read metro7_3
	recordsLine72, err := l.readInput("7_3")
	if err != nil {
		return err
	}
	counter = 0
	for _, record := range recordsLine72 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "7_3")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "7_3")
		}

		if counter == 0 {
			station.AddNeighbour(fork)
			fork.AddNeighbour(station)
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	return nil
}

func (l *FileLoader) loadLine13(n *domain.Network) error {
	// Read metro13_3
	recordsLine133, err := l.readInput("13_3")
	if err != nil {
		return err
	}

	var (
		fork     *domain.Station
		previous *domain.Station
	)
	counter := 0
	for _, record := range recordsLine133 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "13_3")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "13_3")
		}

		if counter == 0 {
			fork = station
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	// Read metro13_1
	recordsLine131, err := l.readInput("13_1")
	if err != nil {
		return err
	}
	counter = 0

	for _, record := range recordsLine131 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "13_1")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "13_1")
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	previous.AddNeighbour(fork)
	fork.AddNeighbour(previous)

	// Read metro7_2
	recordsLine132, err := l.readInput("13_2")
	if err != nil {
		return err
	}
	counter = 0
	for _, record := range recordsLine132 {
		var station *domain.Station
		if s, ok := n.Stations[record]; !ok {
			station = &domain.Station{
				Name: record,
			}
			station.Lines = append(station.Lines, "7_2")
			n.Stations[record] = station
		} else {
			station = s
			station.Lines = append(station.Lines, "7_2")
		}

		if counter > 0 {
			station.AddNeighbour(previous)
			previous.AddNeighbour(station)
		}

		previous = station
		counter++
	}

	previous.AddNeighbour(fork)
	fork.AddNeighbour(previous)

	return nil
}
