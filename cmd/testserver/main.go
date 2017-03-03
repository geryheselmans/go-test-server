package main

import (
	"github.com/geryheselmans/go-test-server/repository"
	"github.com/geryheselmans/go-test-server/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	h := web.New(inMemoryAuthorRepository)

	go h.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		os.Exit(0)
	case <-h.ErrChan():
		os.Exit(0)
	}
}
