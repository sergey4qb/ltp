package app

import (
	"github.com/sergey4qb/ltp/internal/core/ports"
	"github.com/sergey4qb/ltp/internal/infrastructure/external/http/kraken_client"
)

type infrastructure struct {
	tickerClient ports.TickerClient
}

func createInfrastructure() *infrastructure {
	return &infrastructure{
		tickerClient: kraken_client.NewPublicClient(),
	}
}
