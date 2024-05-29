package main

import (
	"context"
	"github.com/sergey4qb/ltp/internal/cmd/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.Create()
	application.Start()

	<-stop
	log.Println("Shutting down server...")
}
