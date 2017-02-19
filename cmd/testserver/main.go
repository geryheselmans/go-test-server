package main

import "github.com/geryheselmans/go-test-server/web"

func main() {
	h := web.New()

	h.Run()
}
