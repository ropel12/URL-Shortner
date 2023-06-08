package main

import (
	"github.com/ropel12/URL-Shortner/handlers"
	"github.com/ropel12/URL-Shortner/repository"
	"github.com/ropel12/URL-Shortner/router"
)

func main() {
	repo := repository.NewMapRepository()
	h := handlers.NewURLHandler(&repo)
	r := router.NewRouter(h)
	r.Run(":8000")
}
