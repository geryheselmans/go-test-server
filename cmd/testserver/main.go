package main

import (
	"github.com/geryheselmans/go-test-server/repository"
	"github.com/geryheselmans/go-test-server/web"
)

func main() {
	inMemoryAuthorRepository := repository.NewInMemoryAuthorRepository()

	h := web.New(inMemoryAuthorRepository)

	h.Run()
}
