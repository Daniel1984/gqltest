package main

import (
	"log"

	"github.com/gqltest/cmd/server/config"
	"github.com/gqltest/cmd/server/router"
	"github.com/gqltest/internals/graceful_shutdown"
	"github.com/gqltest/internals/server"
	"github.com/machinebox/graphql"
)

func main() {
	cfg := config.New()
	gqlClient := graphql.NewClient(cfg.UniswapAPI)
	srv := server.
		New().
		WithAddr(cfg.GetApiPort()).
		WithRouter(router.New(gqlClient))

	go func() {
		log.Printf("starting server at %s", cfg.GetApiPort())
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	graceful_shutdown.Init(func() {
		if err := srv.Close(); err != nil {
			log.Println(err)
		}
	})
}
