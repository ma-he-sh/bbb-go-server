package main

import (
	"net/http"
    "time"
    "log"

	"github.com/gorilla/mux"
    
    routes "github.com/devmarka/bbb-go-server/rest"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

    mainRouter := r.Host("localhost").Subrouter()
    routes.Routes( mainRouter )

    fs := http.FileServer( http.Dir("./public"))
    r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

    server := &http.Server{
        Addr: ":8080",
        Handler: r,
        ReadTimeout: 15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }

	log.Fatal(server.ListenAndServe())
}
