package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/exp/slog"

	server "wb_l2/internal/server/eventserver"
	"wb_l2/internal/service/eventservice"
	"wb_l2/internal/storage/eventstorage"
)

type configuration struct {
	ServerPort int `envconfig:"SERVER_PORT" required:"true"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	storage := eventstorage.NewStorage()
	svc := eventservice.NewService(storage)
	server := server.NewServer(svc, cfg.ServerPort, logger)
	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
}
