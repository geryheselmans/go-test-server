package main

import (
	"github.com/geryheselmans/go-test-server/repository"
	"github.com/geryheselmans/go-test-server/web"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	log.Info("Starting service")

	h := web.New(inMemoryAuthorRepository)

	go h.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		log.Info("Received interrupt, start shutdown procedure")
	case err := <-h.ErrChan():
		log.WithError(err).Error("Error while starting service")
	}

	log.Info("Goodbye")
}
