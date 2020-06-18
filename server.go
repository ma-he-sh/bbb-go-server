package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	db "github.com/devmarka/bbb-go-server/core/db"
	routes "github.com/devmarka/bbb-go-server/routes"
	env "github.com/devmarka/bbb-go-server/env"
)

func main() {
	db.CreateTable()

	r := mux.NewRouter().StrictSlash(true)

	mainRouter := r.Host(env.APPDomain()).Subrouter()
	routes.Routes(mainRouter)
	routes.Rest(mainRouter)

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
