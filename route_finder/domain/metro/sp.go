package metro

import (
	"fmt"
	"log"
	"math"
)

type shortestPathSolver struct {
	network           *Network
	distances         map[*Station]int
	knownNodes        map[*Station]bool
	visitedNodes      map[*Station]bool
	shortestDistances map[*Station]int
	previous          map[*Station]*Station
}

func NewShortestPathSolver(network *Network) *shortestPathSolver {
	return &shortestPathSolver{
		network:           network,
		distances:         make(map[*Station]int),
		knownNodes:        make(map[*Station]bool),
		visitedNodes:      make(map[*Station]bool),
		shortestDistances: make(map[*Station]int),
		previous:          make(map[*Station]*Station),
	}
}

func (sps *shortestPathSolver) initDistances(s *Station) {
	sps.distances[s] = 0
	for _, node := range sps.network.Stations {
		if node != s {
			sps.distances[node] = math.MaxInt32
		}
	}
	sps.knownNodes[s] = true
}

func (sps *shortestPathSolver) findClosestUnvisitedKnownVertex() *Station {
	var closest *Station
	min := math.MaxInt32
	for v := range sps.knownNodes {
		if sps.distances[v] < min && !sps.visitedNodes[v] {
			min = sps.distances[v]
			closest = v
		}
	}
	return closest
}

func (sps *shortestPathSolver) discoverNeighbours(current *Station) {
	for _, v := range current.Neighbours {
		if _, ok := sps.knownNodes[v]; !ok {
			sps.knownNodes[v] = true
		}
	}
}

func (sps *shortestPathSolver) updateDistances(current *Station) {
	for node := range sps.knownNodes {
		if _, ok := sps.visitedNodes[node]; !ok {
			if sps.shortestDistances[current]+current.DistToNeighbours[node] < sps.distances[node] {
				sps.distances[node] = sps.shortestDistances[current] + current.DistToNeighbours[node]
				sps.shortestDistances[node] = sps.shortestDistances[current] + current.DistToNeighbours[node]
				sps.previous[node] = current
			}
		}
	}
}

func (sps *shortestPathSolver) discoverAndUpdate(current *Station) {
	for v, dist := range current.DistToNeighbours {
		if _, ok := sps.knownNodes[v]; !ok {
			sps.knownNodes[v] = true
			if sps.shortestDistances[current]+dist < sps.distances[v] {
				sps.distances[v] = sps.shortestDistances[current] + dist
				sps.shortestDistances[v] = sps.shortestDistances[current] + dist
				sps.previous[v] = current
			}
		}
	}
}

func (sps *shortestPathSolver) Solve(source, dest *Station) []*Station {
	sps.initDistances(source)

	var current *Station
	for {
		if len(sps.visitedNodes) == len(sps.network.Stations) {
			log.Println("all Nodes have been visited")
			break
		}

		current = sps.findClosestUnvisitedKnownVertex()
		sps.discoverAndUpdate(current)
		sps.visitedNodes[current] = true
	}

	c := dest
	path := fmt.Sprintf("%s", c.Name)
	for {
		if c == source {
			break
		}
		c = sps.previous[c]
		path = fmt.Sprintf("%s -> ", c.Name) + path
	}

	log.Println(path)
	log.Println(sps.distances[dest])

	return nil
}
