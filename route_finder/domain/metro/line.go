package metro

type Line struct {
	Name     string
	Number   int
	Stations []*Station
}

func NewLine(name string, number int, stations []*Station) *Line {
	return &Line{
		Name:     name,
		Number:   number,
		Stations: stations,
	}
}
