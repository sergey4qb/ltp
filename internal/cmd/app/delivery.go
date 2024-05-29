package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/ltp/internal/core/delivery/http/router"
	"log"
	"os"
)

type delivery struct {
	http *router.Router
}

func createDelivery(services *services) *delivery {
	gin.SetMode(os.Getenv("GIN_MODE"))
	g := gin.Default()
	r := router.NewRouter(
		g,
		services.priceService,
	)
	return &delivery{
		http: r,
	}
}

func (d *delivery) Start() {
	go func() {
		err := d.http.Run(":" + os.Getenv("HTTP_PORT"))
		if err != nil {
			log.Println("Error starting server:", err)
		}
	}()
}
