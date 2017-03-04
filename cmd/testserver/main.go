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
	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	logger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	logger.Info("Starting service")

	h := web.New(logger, inMemoryAuthorRepository)

	go h.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		logger.Info("Received stop signal, start shutting down")
	case err := <-h.ErrChan():
		logger.Error("Error while starting service", zap.Error(err))
	}

	logger.Info("Goodbye")
}
