package metro

import "github.com/cockroachdb/errors"

type Network struct {
	Stations map[string]*Station
	Lines    []*Line
}

func NewNetwork() *Network {
	return &Network{
		Stations: make(map[string]*Station),
	}
}

func (n *Network) GetStation(name string) (*Station, error) {
	v, ok := n.Stations[name]
	if !ok {
		return nil, errors.Newf("no station %q", name)
	}
	return v, nil
}
