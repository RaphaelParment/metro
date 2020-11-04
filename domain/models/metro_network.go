package models

type MetroStation struct {
	// name of the station
	Name string
	// which metro lines transit via the station
	Lines []int
	// list of possible next stations
	Next map[int]*MetroStation
	// list of possible previous stations
	Prev map[int]*MetroStation
}

func NewMetroStations(name string, line int) *MetroStation {
	var station MetroStation
	station.Name = name
	station.Lines = append(station.Lines, line)

	station.Next = make(map[int]*MetroStation)
	station.Prev = make(map[int]*MetroStation)

	return &station
}

type MetroLine struct {
	Name     string
	Number   int
	Stations []*MetroStation
}

type MetroNetwork struct {
	Lines []*MetroLine
}
