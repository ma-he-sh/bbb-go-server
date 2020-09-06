package api

import (
	"errors"
	"log"
	"time"

	db "github.com/devmarka/bbb-go-server/core/db"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type AuthStruct struct {
	ModeratorPW string `rethinkdb:"moderatorpw"`
	AttendeePW  string `rethinkdb:"attendeepw"`
}

type EventStruct struct {
	Id          string     `rethinkdb:"id"`
	EventName   string     `rethinkdb:"name"`
	EventTime   string     `rethinkdb:"eventtime"`
	Stamp       string     `rethinkdb:"stamp"`
	Auth        AuthStruct `rethinkdb:"auth"`
	Record      bool       `rethinkdb:"record"`
	ToggleEmail bool       `rethinkdb:"show_email"`
	Active      bool       `rethinkdb:"active"`
}

func CreateEvent(event EventStruct, isNew bool) EventStruct {
	setevent := event

	if isNew {
		datetime := time.Now()
		setevent.Stamp = datetime.Format("2016-01-02 15:04:05")
	}

	return setevent
}

func (event *EventStruct) GetEventName() string {
	return event.EventName
}

func (event *EventStruct) SetEventName(name string) {
	event.EventName = name
}

func (event *EventStruct) GetModeratorPW() string {
	return event.Auth.ModeratorPW
}

func (event *EventStruct) SetModeratorPW(passw string) {
	event.Auth.ModeratorPW = passw
}

func (event *EventStruct) GetAttendeePW() string {
	return event.Auth.AttendeePW
}

func (event *EventStruct) SetAttendeePW(passw string) {
	event.Auth.AttendeePW = passw
}

func (event *EventStruct) AllowRecord() bool {
	return event.Record
}

func (event *EventStruct) ShowEmail() bool {
	return event.ToggleEmail
}

// EventList
func EventList() ([]EventStruct, error) {
	res, err := rethinkdb.Table(db.TBEvent).Run(db.Session)
	var events []EventStruct
	if err != nil {
		return events, err
	}
	res.All(&events)
	defer res.Close()
	return events, err
}

// Get recordings
func RecordList() ([]map[string]interface{}, error) {
	var records []map[string]interface{}

	var recordIDs = []string{}
	eventRecords := BBBEventRecordings(recordIDs)
	for _, recordObj := range eventRecords.Recording {
		targetURL := ""
		for _, formatObj := range recordObj.Playback.Format {
			if formatObj.Type == "presentation" {
				targetURL = formatObj.Url
			}
		}
		record := map[string]interface{}{
			"meetingid":           recordObj.MeetingID,
			"meetingName":         recordObj.Name,
			"meetingParticipants": recordObj.Participants,
			"publish":             recordObj.Published,
			"url":                 targetURL,
		}
		records = append(records, record)
	}
	return records, nil
}

// EventExists :: event exists
func EventExists(eventID string) bool {
	var count int
	var exists bool = false
	res, err := rethinkdb.Table(db.TBEvent).GetAll(eventID).Count().Run(db.Session)
	if err != nil {
		log.Fatal(err)
	}

	res.One(&count)
	defer res.Close()
	if count == 1 {
		exists = true
	}
	return exists
}

// CreateEvent :: create new event
func InsertEvent(event EventStruct) (bool, error) {
	res, err := rethinkdb.Table(db.TBEvent).Insert(event).RunWrite(db.Session)
	if res.Inserted == 1 {
		return true, err
	}
	return false, err
}

// get event :: get event data
func GetEvent(eventID string) (EventStruct, error) {
	var event EventStruct
	found := EventExists(eventID)
	if !found {
		return event, errors.New("event_not_found")
	}

	res, err := rethinkdb.Table(db.TBEvent).Get(eventID).Run(db.Session)
	if err != nil {
		return event, errors.New("event_not_found")
	}
	res.One(&event)
	defer res.Close()

	return event, err
}

// Delete Event
func DeleteEvent(eventID string) (bool, error) {
	found := EventExists(eventID)
	if !found {
		return true, errors.New("event_not_found")
	}
	res, err := rethinkdb.Table(db.TBEvent).Get(eventID).Delete().RunWrite(db.Session)
	if err != nil {
		return false, errors.New("server_error")
	}
	if res.Deleted == 1 {
		return true, nil
	}
	return false, errors.New("server_error")
}
