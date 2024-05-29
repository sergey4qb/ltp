package ports

import (
	"context"
	"github.com/sergey4qb/ltp/internal/dto"
)

type TickerClient interface {
	GetTickerInfo(ctx context.Context, symbol string) (*dto.TickerResponse, error)
}
