package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"nody/controller"
	"nody/model"

	"github.com/go-chi/chi/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateUser(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			panic(err)
		}
		if err := controller.CreateUser(ctx, driver, user); err != nil {
			panic(err)
		}
		w.Write([]byte("User created"))
	}
}

func CreateIPv4(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ipv4 model.IPv4
		if err := json.NewDecoder(r.Body).Decode(&ipv4); err != nil {
			panic(err)
		}
		if err := controller.CreateIPv4(ctx, driver, ipv4); err != nil {
			panic(err)
		}
		w.Write([]byte("IP created"))
	}
}

func CreateHasIP(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hasip model.HasIP
		if err := json.NewDecoder(r.Body).Decode(&hasip); err != nil {
			panic(err)
		}
		if err := controller.CreateHasIP(ctx, driver, hasip); err != nil {
			panic(err)
		}
		w.Write([]byte("Relationship created."))
	}
}

func GetUsernames(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		names, err := controller.GetUsernames(ctx, driver)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(names)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
}

func GetIPsByUsername(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		ipv4s, err := controller.GetIPsByUsername(ctx, driver, name)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(ipv4s)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
}

func GetUsersByIP(ctx context.Context, driver neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := chi.URLParam(r, "ip")
		users, err := controller.GetUsersByIP(ctx, driver, ip)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
}
