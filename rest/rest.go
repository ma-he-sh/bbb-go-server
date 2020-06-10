package rest

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Routes(routes *mux.Router) {
    routes.HandleFunc("/", rootRoute)
}

func rootRoute(w http.ResponseWriter, r *http.Request ) {
    fmt.Fprintf(w, "ROOT")
}
