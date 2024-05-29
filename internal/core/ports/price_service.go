package ports

import (
	"context"
	"github.com/sergey4qb/ltp/internal/core/domain/valueobject"
)

type PriceService interface {
	GetPricesForPairs(ctx context.Context, pairs []*valueobject.Pair) ([]valueobject.LTP, error)
}
