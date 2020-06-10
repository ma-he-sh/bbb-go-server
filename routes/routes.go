package rest

import (
    "net/http"

    "github.com/gorilla/mux"

    templ "github.com/devmarka/bbb-go-server/template"
)

func Routes(routes *mux.Router) {
    routes.HandleFunc("/", rootRoute)
}

func rootRoute(w http.ResponseWriter, r *http.Request ) {
    page := templ.PageObj("Home")
    page.SetBody("<div>HELLO</div>")
    page.IsAdmin(false)
    page.SetFRScripts("admin.js")
    templ.Render(w, "app", page.GetTemplPayload())
}

func Throw400(w http.ResponseWriter, r *http.Request ) {
    w.WriteHeader(404)
    w.Write([]byte(`Error Page Not Found`))
    return
}
