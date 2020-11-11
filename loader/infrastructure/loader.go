package infrastructure

import (
	"context"
	"github.com/RaphaelParment/metro/loader/domain"
)

type Loader interface {
	Load(ctx context.Context, n *domain.Network, number int) error
}
