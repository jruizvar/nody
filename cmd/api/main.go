package main

import (
	"context"
	"encoding/json"
	"net/http"
	"nody/controler"
	"nody/db"
	"nody/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	// neo4j driver
	driver := db.Driver()

	ctx := context.Background()
	defer driver.Close(ctx)

	// fill the graph
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			panic(err)
		}
		if err := controler.CreateUser(ctx, driver, user); err != nil {
			panic(err)
		}
		w.Write([]byte("User created"))
	})

	r.Post("/ipv4s", func(w http.ResponseWriter, r *http.Request) {
		var ipv4 model.IPv4
		if err := json.NewDecoder(r.Body).Decode(&ipv4); err != nil {
			panic(err)
		}
		if err := controler.CreateIPv4(ctx, driver, ipv4); err != nil {
			panic(err)
		}
		w.Write([]byte("IP created"))
	})

	r.Post("/hasip", func(w http.ResponseWriter, r *http.Request) {
		var hasip model.HasIP
		if err := json.NewDecoder(r.Body).Decode(&hasip); err != nil {
			panic(err)
		}
		if err := controler.CreateHasIP(ctx, driver, hasip); err != nil {
			panic(err)
		}
		w.Write([]byte("Relationship HAS_IP created."))
	})

	// query the graph
	r.Get("/users/names", func(w http.ResponseWriter, r *http.Request) {
		names, err := controler.GetUsernames(ctx, driver)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(names)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	})

	r.Get("/users/{name}/ipv4s", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		ipv4s, err := controler.GetIPsByUsername(ctx, driver, name)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(ipv4s)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	})

	r.Get("/ipv4s/{ip}/users", func(w http.ResponseWriter, r *http.Request) {
		ip := chi.URLParam(r, "ip")
		users, err := controler.GetUsersByIP(ctx, driver, ip)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	})

	http.ListenAndServe(":3000", r)
}
