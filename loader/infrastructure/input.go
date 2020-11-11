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

func (l *FileLoader) Load(ctx context.Context, n *domain.Network, number int) error {
	records, err := l.readInput(number)
	if err != nil {
		return err
	}
	return l.loadStraightLine(n, records, number)
}

func (l *FileLoader) readInput(number int) ([]string, error) {
	path := fmt.Sprintf("../data/metro_%d.txt", number)
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

func (l *FileLoader) loadStraightLine(n *domain.Network, stations []string, number int) error {
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
			station.Neighbours = append(station.Neighbours, previous)
			previous.Neighbours = append(previous.Neighbours, station)
		}

		previous = station
		counter++
	}
	return nil
}
