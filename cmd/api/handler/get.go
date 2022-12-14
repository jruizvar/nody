package handler

import (
	"encoding/json"
	"net/http"
	"nody/controler"
	"nody/db"

	"github.com/go-chi/chi/v5"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func GetUsernames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = db.Driver(ctx)
	names, err := controler.GetUsernames(ctx)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(names)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func GetIPsByUsername(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	ctx := r.Context()
	ctx = db.Driver(ctx)
	ipv4s, err := controler.GetIPsByUsername(ctx, name)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(ipv4s)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func GetUsersByIP(w http.ResponseWriter, r *http.Request) {
	ip := chi.URLParam(r, "ip")
	ctx := r.Context()
	ctx = db.Driver(ctx)
	users, err := controler.GetUsersByIP(ctx, ip)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}
