package main

import (
	"log"
	"net/http"
	"time"
	"crypto/tls"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"

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

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}

	server := &http.Server{
		Addr:         ":443",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Fatal(server.ListenAndServe())
	go http.ListenAndServe(":8080", certManager.HTTPHandler(nil))
	server.ListenAndServeTLS("", "")
}
