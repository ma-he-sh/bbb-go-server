package rest

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	api "github.com/devmarka/bbb-go-server/core/api"
	session "github.com/devmarka/bbb-go-server/core/session"
	env "github.com/devmarka/bbb-go-server/env"
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
	routes.HandleFunc("/admin/delete/event", deleteEvent).Methods("POST")
	routes.HandleFunc("/admin/gen/joinlink", genJoinLink).Methods("POST")
}

func getEventID(eventname string) string {
	now := time.Now()
	stamp := now.Unix()
	return strings.ToLower(strings.ReplaceAll(eventname, " ", "_")) + "_" + strconv.FormatInt(stamp, 10)
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
		Id:        getEventID(dataForm["eventName"].(string)),
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

func deleteEvent(w http.ResponseWriter, r *http.Request) {
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

	deleted, err := api.DeleteEvent(dataForm["eventid"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if deleted {
		send["success"] = true
	}
	JSONPayload(w, send)
	return
}

func genJoinLink(w http.ResponseWriter, r *http.Request) {
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
		"payload": map[string]string{
			"url": "",
		},
	}

	event, err := api.GetEvent(dataForm["eventid"].(string))
	if err != nil {
		JSONPayload(w, send)
		return
	}

	params := url.Values{}
	params.Add("uname", dataForm["username"].(string))
	params.Add("code", event.GetAttendeePW())

	joinURL := env.APPDOMAIN_name() + "/join/" + event.Id + "/auth?" + params.Encode()

	send["success"] = true
	send["payload"] = map[string]string{
		"url": joinURL,
	}

	JSONPayload(w, send)
	return
}
