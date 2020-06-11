package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	templ "github.com/devmarka/bbb-go-server/template"
)

func Routes(routes *mux.Router) {
	routes.HandleFunc("/", rootRoute)
	routes.HandleFunc("/admin/login", adminLogin).Methods("GET", "POST")
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	page := templ.PageObj("Home")
	page.SetBody("<div>HELLO</div>")
	page.IsAdmin(false)
	page.SetFRScripts("frontend.js")
	templ.Render(w, "app", page.GetTemplPayload())
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	page := templ.PageObj("Signin")
	page.IsAdmin(false)
	page.SetFRScripts("frontend.js")
	if r.Method == http.MethodGet {
		page.SetBody(templ.AdminLoginForm())
		templ.Render(w, "app", page.GetTemplPayload())
		return
	} else if r.Method == http.MethodPost {
		page.SetBody(templ.AdminLoginForm())
		templ.Render(w, "app", page.GetTemplPayload())
		return
	} else {
		Throw400(w, r)
	}
}

func Throw400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte(`Error Page Not Found`))
	return
}
