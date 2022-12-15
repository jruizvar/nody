package main

import (
	"context"
	"net/http"
	"nody/db"
	"nody/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	// never-cancelling context
	ctx := context.Background()

	// neo4j driver
	driver := db.Driver()
	defer driver.Close(ctx)

	// fill the graph
	r.Post("/users", handler.CreateUser(ctx, driver))
	r.Post("/ipv4s", handler.CreateIPv4(ctx, driver))
	r.Post("/hasip", handler.CreateHasIP(ctx, driver))

	// query the graph
	r.Get("/users/names", handler.GetUsernames(ctx, driver))
	r.Get("/users/{name}/ipv4s", handler.GetIPsByUsername(ctx, driver))
	r.Get("/ipv4s/{ip}/users", handler.GetUsersByIP(ctx, driver))

	http.ListenAndServe(":3000", r)
}
