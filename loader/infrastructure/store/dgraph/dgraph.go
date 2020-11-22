package dgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RaphaelParment/metro/loader/domain"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"log"
)

type Station struct {
	UID        string   `json:"uid,omitempty"`
	Name       string   `json:"name,omitempty"`
	Lines      []string `json:"lines,omitempty"`
	Neighbours []string `json:"neighbours,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type CancelFunc func()

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func Load(network *domain.Network) error {
	dg, cancel := getDgraphClient()
	defer cancel()

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		lines: [string] @index(hash) .
		type: string .
		neighbours: [uid] .
		type Station {
			name: string
			lines: [string]
			neighbours: [Station]
		}`

	ctx := context.Background()
	if err := dg.Alter(ctx, op); err != nil {
		return err
	}

	stations := make([]Station, len(network.Stations))
	counter := 0
	for _, station := range network.Stations {
		s := Station{
			UID:   "_:" + station.Name,
			Name:  station.Name,
			Lines: station.Lines,
		}
		stations[counter] = s
		counter++
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(stations)
	if err != nil {
		return err
	}
	mu.SetJson = pb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return err
	}

	for _, s := range network.Stations {
		for _, neighbour := range s.Neighbours {

			m1 := fmt.Sprintf("<%s> <neighbours> <%s> .", response.Uids[s.Name], response.Uids[neighbour.Name])
			mu := &api.Mutation{
				SetNquads: []byte(m1),
				CommitNow: true,
			}
			if _, err := dg.NewTxn().Mutate(ctx, mu); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
