package store

import (
	"context"
	"github.com/RaphaelParment/metro/loader/domain"
)

type Store interface {
	Save(ctx context.Context, network *domain.Network) error
}
