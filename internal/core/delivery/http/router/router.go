package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/ltp/internal/core/delivery/http/controllers"
	"github.com/sergey4qb/ltp/internal/core/ports"
	"time"
)

type Router struct {
	gin            *gin.Engine
	PriceContoller *controllers.PriceController
}

func NewRouter(
	gin *gin.Engine,
	priceService ports.PriceService,
) *Router {
	r := &Router{
		gin: gin,
	}
	r.PriceContoller = controllers.NewPriceController(priceService)
	r.setupCORS()
	r.setupRoutes()
	return r
}

func (r *Router) Run(port string) error {
	return r.gin.Run(port)
}

func (r *Router) setupCORS() {
	r.gin.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://localhost:8080", "http://localhost:8080", "https://1c08-31-128-77-167.ngrok-free.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "https://95a0-31-128-77-167.ngrok-free.app" || origin == "https://localhost:8080" || origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func (r *Router) setupRoutes() {
	r.setupApiRoutes()
}

func (r *Router) setupApiRoutes() {
	group := r.gin.Group("/api/v1")
	group.GET("/ltp", r.PriceContoller.GetPricesForPairs)
}
