package metro

type Station struct {
	// name of the station
	Name string
	// which metro lines transit via the station
	Lines []int

	Neighbours       []*Station
	DistToNeighbours map[*Station]int
}

func NewStation(name string, line int) *Station {
	var station Station
	station.Name = name
	station.Lines = append(station.Lines, line)

	// all distances are 0
	station.DistToNeighbours = make(map[*Station]int)
	return &station
}
