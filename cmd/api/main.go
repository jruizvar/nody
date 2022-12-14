package main

import (
	"net/http"
	"nody/cmd/api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	// fill the graph
	r.Post("/users", handler.CreateUser)
	r.Post("/ipv4s", handler.CreateIPv4)
	r.Post("/hasip", handler.CreateHasIP)

	// query the graph
	r.Get("/users/names", handler.GetUsernames)
	r.Get("/users/{name}/ipv4s", handler.GetIPsByUsername)
	r.Get("/ipv4s/{ip}/users", handler.GetUsersByIP)

	http.ListenAndServe(":3000", r)
}
