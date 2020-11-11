package domain

import (
	"log"
	"strconv"
	"strings"
)

type Station struct {
	Position      []float64
	IDRefZdl      string
	UniqueCode    string
	Name          string
	Label         string
	IDRefLiga     string
	IDRefLigc     string
	Lines         string
	TransportType string
	Fer           bool
	Train         bool
	Rer           bool
	Metro         bool
	Tramway       bool
	Shuttle       bool
	Val           bool
	Terfer        bool
	Tertrain      bool
	Terrer        bool
	Termetro      bool
	Tertram       bool
	Ternavette    bool
	Terval        bool
	Operator      string
	Principal     bool
	Idf           bool
	X             float64
	Y             float64
}

func NewStation(csvLine []string) *Station {
	coords := strings.Split(csvLine[0], ",")
	log.Println(coords)
	long, _ := strconv.ParseFloat(coords[0], 64)
	lat, _ := strconv.ParseFloat(coords[1], 64)

	fer, _ := strconv.ParseBool(csvLine[9])
	train, _ := strconv.ParseBool(csvLine[10])
	rer, _ := strconv.ParseBool(csvLine[11])
	metro, _ := strconv.ParseBool(csvLine[12])
	tramway, _ := strconv.ParseBool(csvLine[13])
	shuttle, _ := strconv.ParseBool(csvLine[14])
	val, _ := strconv.ParseBool(csvLine[15])
	terfer, _ := strconv.ParseBool(csvLine[16])
	tertrain, _ := strconv.ParseBool(csvLine[17])
	terrer, _ := strconv.ParseBool(csvLine[18])
	termetro, _ := strconv.ParseBool(csvLine[19])
	tertram, _ := strconv.ParseBool(csvLine[20])
	ternavette, _ := strconv.ParseBool(csvLine[21])
	terval, _ := strconv.ParseBool(csvLine[22])

	principal, _ := strconv.ParseBool(csvLine[24])
	idf, _ := strconv.ParseBool(csvLine[25])

	x, _ := strconv.ParseFloat(csvLine[26], 64)
	y, _ := strconv.ParseFloat(csvLine[27], 64)
	s := Station{
		Position:      []float64{long, lat},
		IDRefZdl:      csvLine[1],
		UniqueCode:    csvLine[2],
		Name:          csvLine[3],
		Label:         csvLine[4],
		IDRefLiga:     csvLine[5],
		IDRefLigc:     csvLine[6],
		Lines:         csvLine[7],
		TransportType: csvLine[8],
		Fer:           fer,
		Train:         train,
		Rer:           rer,
		Metro:         metro,
		Tramway:       tramway,
		Shuttle:       shuttle,
		Val:           val,
		Terfer:        terfer,
		Tertrain:      tertrain,
		Terrer:        terrer,
		Termetro:      termetro,
		Tertram:       tertram,
		Ternavette:    ternavette,
		Terval:        terval,
		Operator:      csvLine[23],
		Principal:     principal,
		Idf:           idf,
		X:             x,
		Y:             y,
	}
	return &s
}
