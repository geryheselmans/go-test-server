package main

import (
	"github.com/geryheselmans/go-test-server/repository"
	"github.com/geryheselmans/go-test-server/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	log.Info("Starting service")

	h := web.New(inMemoryAuthorRepository)

	h.Run()

	log.Info("Stop service")
}
