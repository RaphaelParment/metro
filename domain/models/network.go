package models

import (
	"encoding/csv"
	"io"
	"log"
)

type Network struct {
	Stations []*Station
}

func NewNetwork(input *csv.Reader) (*Network, error) {
	var stations []*Station
	for {
		record, err := input.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read line; %v", err)
			return nil, err
		}

		s := NewStation(record)

		stations = append(stations, s)
	}

	n := Network{
		Stations: stations,
	}

	return &n, nil
}
