package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	templ "github.com/devmarka/bbb-go-server/template"
	session "github.com/devmarka/bbb-go-server/core/session"
	env "github.com/devmarka/bbb-go-server/env"
)

func Routes(routes *mux.Router) {
	routes.HandleFunc("/", rootRoute)
	routes.HandleFunc("/admin/login", adminLogin).Methods("GET", "POST")
	routes.HandleFunc("/admin/dashboard", adminDashboard).Methods("GET")
	routes.HandleFunc("/admin/signout", adminSignout).Methods("GET")
	routes.HandleFunc("/event/{eventid}", eventHandle).Methods("GET")
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	session.SessionRedirect(w, r, "/admin/dashboard")

	page := templ.PageObj("Home")
	page.SetBody("<div>HELLO</div>")
	page.IsAdmin(false)
	page.SetFRScripts("frontend.js")
	templ.Render(w, "app", page.GetTemplPayload())
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	session.SessionRedirect(w, r, "/admin/dashboard")

	page := templ.PageObj("Signin")
	page.IsAdmin(false)
	page.SetFRScripts("frontend.js")
	if r.Method == http.MethodGet {
		page.SetBody(templ.AdminLoginForm(""))
		templ.Render(w, "app", page.GetTemplPayload())
		return
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		strEmail := r.FormValue("str_email")
		strPassw := r.FormValue("str_passw")

		if strEmail == env.APPLogin() && strPassw == env.APPPassw() {
			sess, _ := session.CookieStore.Get(r, "user_cookie")
			sess.Values["user"] = "admin"
			sess.Values["authed"] = true
			err := sess.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/admin/login?err=autherr", http.StatusSeeOther)
		return
	} else {
		Throw400(w, r)
	}
}

func eventHandle(w http.ResponseWriter, r *http.Request) {
	page := templ.PageObj("Event")
	page.SetBody("<div>HELLO</div>")
	page.IsAdmin(false)
	page.SetFRScripts("frontend.js")
	templ.Render(w, "app", page.GetTemplPayload())
}

func adminDashboard(w http.ResponseWriter, r *http.Request) {
	session.SessionAuthCheck(w, r, "/")

	page := templ.PageObj("Dashboard")
	page.SetBody("<div>DASHBOARD</div>")
	page.IsAdmin(true)
	page.SetBKScripts("backend.js")
	templ.Render(w, "app", page.GetTemplPayload())
}

func adminSignout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.CookieStore.Get(r, "user_cookie")

	sess.Values["user"] = ""
	sess.Values["authed"] = false
	sess.Options.MaxAge = -1
	err  := sess.Save(r, w)
	if err  != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Throw400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte(`Error Page Not Found`))
	return
}
