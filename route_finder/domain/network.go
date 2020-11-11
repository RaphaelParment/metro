package domain

import (
	"encoding/csv"
	"io"
	"log"
)

type TransportNetwork struct {
	Stations []*Station
}

func NewTransportNetwork(input *csv.Reader) (*TransportNetwork, error) {
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

	n := TransportNetwork{
		Stations: stations,
	}

	return &n, nil
}
