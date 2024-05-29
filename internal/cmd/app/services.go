package app

import (
	"github.com/sergey4qb/ltp/internal/core/ports"
	"github.com/sergey4qb/ltp/internal/core/service/price_service"
)

type services struct {
	priceService ports.PriceService
}

func createServices(infrastructure *infrastructure) *services {
	priceService := price_service.NewService(infrastructure.tickerClient)

	return &services{
		priceService: priceService,
	}
}
