package main

import (
	"github.com/geryheselmans/go-test-server/repository"
	"github.com/geryheselmans/go-test-server/web"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logConfig := os.Getenv("LOGCONFIG")

	var logger *zap.Logger
	var err error

	if logConfig == "PROD" {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
	}

	zap.ReplaceGlobals(logger)

	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	h := web.New(logger, inMemoryAuthorRepository)

	logger.Info("Starting service")
	go h.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		logger.Info("Received stop signal")
	case err := <-h.ErrChan():
		logger.Error("Error while starting service", zap.Error(err))
	}

	logger.Info("Goodbye")
}
