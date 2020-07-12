package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	api "github.com/devmarka/bbb-go-server/core/api"
	session "github.com/devmarka/bbb-go-server/core/session"
	env "github.com/devmarka/bbb-go-server/env"
	templ "github.com/devmarka/bbb-go-server/template"
)

func Routes(routes *mux.Router) {
	routes.HandleFunc("/", rootRoute)
	routes.HandleFunc("/admin/login", adminLogin).Methods("GET", "POST")
	routes.HandleFunc("/admin/dashboard/", adminDashboard).Methods("GET")
	routes.HandleFunc("/admin/dashboard/event/add/", createEvent).Methods("GET")
	routes.HandleFunc("/admin/dashboard/event/edit/{eventid}", editEvent).Methods("GET")
	routes.HandleFunc("/admin/dashboard/join/{eventid}/", adminJoin).Methods("GET")
	routes.HandleFunc("/admin/signout", adminSignout).Methods("GET")
	routes.HandleFunc("/event/{eventid}/", eventHandle).Methods("GET", "POST")
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	session.SessionRedirect(w, r, "/admin/dashboard")

	if r.Method == http.MethodGet {
		page := templ.PageObj("Home")
		page.SetBody("<div>HELLO</div>")
		page.IsAdmin(false)
		templ.Render(w, "app", page.GetTemplPayload())
	} else {
		Throw400(w, r)
	}
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	session.SessionRedirect(w, r, "/admin/dashboard")

	page := templ.PageObj("Signin")
	page.IsAdmin(false)
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
	return
}

func adminJoin(w http.ResponseWriter, r *http.Request) {
	session.SessionAuthCheck(w, r, "/")

	params := mux.Vars(r)
	event, err := api.GetEvent(params["eventid"])
	if err != nil {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	// get user session
	sess, _ := session.CookieStore.Get(r, "user_cookie")
	strName := sess.Values["user"].(string)
	strAccessCode := event.GetModeratorPW()
	eventID := params["eventid"]

	logouturl := env.APPDOMAIN_name() + `/admin/dashboard`
	url, allowed := api.BBBJoinMeetingURL(eventID, strName, strAccessCode, logouturl)
	if allowed {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	return
}

func eventHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == http.MethodGet {
		event, err := api.GetEvent(params["eventid"])
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		eventdata := map[string]interface{}{
			"name":         event.EventName,
			"eventid":      params["eventid"],
			"active":       event.Active,
			"time":         event.EventTime,
			"toggle_email": event.ShowEmail(),
		}

		page := templ.PageObj("Event")
		page.IsAdmin(false)
		page.SetBody(templ.ClientLoginForm(eventdata))
		templ.Render(w, "app", page.GetTemplPayload())
		return
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		eventID := r.FormValue("str_eventid")
		//strEmail := r.FormValue("str_email")
		strName := r.FormValue("str_name")
		strAccessCode := r.FormValue("str_token")

		event, err := api.GetEvent(eventID)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		logouturl := r.Host + r.URL.Path
		if strAccessCode == event.GetAttendeePW() || strAccessCode == event.GetModeratorPW() {
			url, allowed := api.BBBJoinMeetingURL(eventID, strName, strAccessCode, logouturl)
			if allowed {
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}
		}
		http.Redirect(w, r, r.URL.Path+"?auth=error", http.StatusSeeOther)
		return
	} else {
		Throw400(w, r)
	}
	return
}

func adminDashboard(w http.ResponseWriter, r *http.Request) {
	session.SessionAuthCheck(w, r, "/")

	if r.Method == http.MethodGet {
		payload := templ.PagePayload{
			Page:      "dashboard",
			EventList: nil,
		}

		tab := r.FormValue("tab")
		if tab != "" {
			payload.Page = tab
		}

		if payload.Page == "dashboard" {
			events, _ := api.EventList()
			var eventRenderList = []map[string]interface{}{}
			for _, event := range events {
				var eventData = map[string]interface{}{
					"eventid":     event.Id,
					"eventName":   event.EventName,
					"eventRecord": event.Record,
					"eventActive": event.Active,
					"eventTime":   event.EventTime,
					"attendeePW":  event.GetAttendeePW(),
					"moderatorPW": event.GetModeratorPW(),
					"domain":      env.APPDOMAIN_name(),
				}
				eventRenderList = append(eventRenderList, eventData)
			}
			payload.EventList = eventRenderList
		}

		page := templ.PageObj("Dashboard")
		page.SetBody(templ.AdminDashboard(payload))
		page.IsAdmin(true)
		templ.Render(w, "app", page.GetTemplPayload())
	} else {
		Throw400(w, r)
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	session.SessionAuthCheck(w, r, "/")

	if r.Method == http.MethodGet {
		payload := templ.PagePayload{
			Page: "create_event",
		}

		page := templ.PageObj("Dashboard::Event")
		page.SetBody(templ.AdminDashboard(payload))
		page.IsAdmin(true)
		templ.Render(w, "app", page.GetTemplPayload())
		return
	} else {
		Throw400(w, r)
	}
}

func editEvent(w http.ResponseWriter, r *http.Request) {
	session.SessionAuthCheck(w, r, "/")

	params := mux.Vars(r)
	if params["eventid"] == "" {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}
}

func adminSignout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.CookieStore.Get(r, "user_cookie")

	sess.Values["user"] = ""
	sess.Values["authed"] = false
	sess.Options.MaxAge = -1
	err := sess.Save(r, w)
	if err != nil {
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
