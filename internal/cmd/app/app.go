package app

type App struct {
	infrastructure *infrastructure
	services       *services
	delivery       *delivery
}

func Create() *App {
	i := createInfrastructure()
	s := createServices(i)
	d := createDelivery(s)
	return &App{
		services:       s,
		infrastructure: i,
		delivery:       d,
	}
}

func (a *App) Start() {
	a.delivery.Start()
}
