package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/ltp/internal/core/domain/valueobject"
	"github.com/sergey4qb/ltp/internal/core/ports"
	"net/http"
)

type PriceController struct {
	service ports.PriceService
}

func NewPriceController(service ports.PriceService) *PriceController {
	return &PriceController{service: service}
}

func (controller *PriceController) GetPricesForPairs(c *gin.Context) {
	pair1 := valueobject.NewPair("BTC", "USD")
	pair2 := valueobject.NewPair("BTC", "CHF")
	pair3 := valueobject.NewPair("BTC", "EUR")
	prices, err := controller.service.GetPricesForPairs(
		c,
		[]*valueobject.Pair{pair1, pair2, pair3},
	)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	response := LTPResponse{
		LTP: prices,
	}
	c.JSON(http.StatusOK, response)
}
