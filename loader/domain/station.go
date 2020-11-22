package domain

type Station struct {
	Name       string
	Lines      []string
	Neighbours []Station
}

func (s *Station) AddNeighbour(neighbour *Station) {
	alreadyNeighbour := false
	for _, n := range s.Neighbours {
		if n.Name == neighbour.Name {
			alreadyNeighbour = true
		}
	}

	if !alreadyNeighbour {
		s.Neighbours = append(s.Neighbours, *neighbour)
	}
}
