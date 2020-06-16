package rest

import (
	"net/http"

	"encoding/json"

	api "github.com/devmarka/bbb-go-server/core/api"
	session "github.com/devmarka/bbb-go-server/core/session"
	"github.com/gorilla/mux"
)

func JSONPayload(w http.ResponseWriter, data interface{}) {
	sendData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(sendData)
}

func Rest(routes *mux.Router) {
	routes.HandleFunc("/admin/create/event", addEvent).Methods("POST")
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	isAdmin := session.SessionAdminCheck(w, r)
	if !isAdmin {
		http.Error(w, "auth error", 500)
		return
	}

	var dataForm = make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&dataForm)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	send := map[string]interface{}{
		"success": false,
	}

	allowRecord := false
	if dataForm["record"].(string) == "_checked" {
		allowRecord = true
	}

	eventActive := false
	if dataForm["active"].(string) == "_checked" {
		eventActive = true
	}

	eventdata := api.CreateEvent(api.EventStruct{
		EventName: dataForm["eventName"].(string),
		EventTime: "",
		Auth: api.AuthStruct{
			ModeratorPW: dataForm["moderatorpw"].(string),
			AttendeePW:  dataForm["attendeepw"].(string),
		},
		Record: allowRecord,
		Active: eventActive,
	}, true)

	added, err := api.InsertEvent(eventdata)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if added {
		send["success"] = true
	}

	JSONPayload(w, send)
	return
}