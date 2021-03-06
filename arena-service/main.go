package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hellodudu/Ultimate/arena-service/arena"
	datastore "github.com/hellodudu/Ultimate/arena-service/db"
	"github.com/hellodudu/Ultimate/utils"
	logger "github.com/hellodudu/Ultimate/utils/log"
	"github.com/micro/go-micro/v2"
	micro_logger "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	log "github.com/rs/zerolog/log"
)

func main() {
	logger.InitLogger("arena-service")

	ds, err := datastore.NewDatastore()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	// New Service
	micro_logger.Init(micro_logger.WithOutput(logger.Logger))
	service := micro.NewService(
		micro.Name("ultimate-service-arena"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		// micro.Version("latest"),
		// micro.Transport(transport.NewTransport()),
	)

	// Initialise service
	service.Init()

	// new arena
	arena, err := arena.NewArena(context.Background(), service, ds)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("ultimate.service.arena", service.Server(), arena.SubHandler)

	// Run service
	go func() {
		defer utils.CaptureException()
		if err := service.Run(); err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	// Run datastore
	go ds.Run()

	// server exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		log.Info().Msgf("ultimate server closing down (signal: %v)", sig)

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			arena.Stop()
			ds.Stop()
			log.Info().Msg("server exit safely")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
