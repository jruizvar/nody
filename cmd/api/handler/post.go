package handler

import (
	"encoding/json"
	"net/http"
	"nody/controler"
	"nody/db"
	"nody/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}
	ctx := r.Context()
	ctx = db.Driver(ctx)
	if err := controler.CreateUser(ctx, user); err != nil {
		panic(err)
	}
	w.Write([]byte("User created"))
}

func CreateIPv4(w http.ResponseWriter, r *http.Request) {
	var ipv4 model.IPv4
	if err := json.NewDecoder(r.Body).Decode(&ipv4); err != nil {
		panic(err)
	}
	ctx := r.Context()
	ctx = db.Driver(ctx)
	if err := controler.CreateIPv4(ctx, ipv4); err != nil {
		panic(err)
	}
	w.Write([]byte("IP created"))
}

func CreateHasIP(w http.ResponseWriter, r *http.Request) {
	var hasip model.HasIP
	if err := json.NewDecoder(r.Body).Decode(&hasip); err != nil {
		panic(err)
	}
	ctx := r.Context()
	ctx = db.Driver(ctx)
	if err := controler.CreateHasIP(ctx, hasip); err != nil {
		panic(err)
	}
	w.Write([]byte("Relationship HAS_IP created."))
}
