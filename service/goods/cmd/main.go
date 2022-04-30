package main

import (
	"context"
	"os"
	"os/signal"
	"price-chart/service/goods/internal/app"
)

func init() {
	app.InitEnv(".env")
	app.InitDB()
	app.Migrate()
}

func main() {
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	done := make(chan struct{})
	defer close(done)

	signal.Notify(sigs, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		app.Run(ctx)
		done <- struct{}{}
	}()

	<-sigs
	cancel()
	<-done
}
