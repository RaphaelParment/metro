package domain

type Station struct {
	Name       string
	Lines      []int
	Neighbours []*Station
}