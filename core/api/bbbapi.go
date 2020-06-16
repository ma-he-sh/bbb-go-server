package api

import (
	"crypto/sha1"
	"encoding/hex"

	env "github.com/devmarka/bbb-go-server/env"
)

func BBBCreateEvent(eventid string, logouturl string) string {

	event, _ := GetEvent(eventid)

	record := "false"
	if event.Record {
		record = "true"
	}

	arrParams := map[string]string{
		"name":            event.EventName,
		"meetingID":       event.Id,
		"attendeePW":      event.GetAttendeePW(),
		"moderatorPW":     event.GetModeratorPW(),
		"logoutURL":       logouturl,
		"record":          record,
		"meta_suuid":      "",
		"welcome":         "",
		"maxParticipants": "30",
		"isBreakout":      "false",
	}

	return BBBBuildURL("create", arrParams)
}

func BBBEndEvent(eventid string) {

}

func BBBGetEventURL(eventid string, username string, password string) string {
	return ""
}

func BBBEventRunning(eventid string) bool {
	return true
}

func BBBEventRecording(eventid string) {

}

func BBBEventRecordings([]string) []string {
	var stringArr = []string{}
	return stringArr
}

func BBBEventDeleteRecording(recorid string) {

}

func BBBBuildURL(requestType string, args map[string]string) string {
	reqURL := env.BBBHost() + "api/" + requestType + "?"
	params := ""
	index := 0
	for k, v := range args {
		if index == 0 {
			params += ""
		} else {
			params += "&"
		}

		params += k + "=" + v
		index += 1
	}

	checksum := sha1.Sum([]byte(requestType + params + env.BBBKey()))
	reqURL = reqURL + params + "&checksum=" + hex.EncodeToString(checksum[:])
	return reqURL
}

func BBBBuildDownloadURL(recordid string) {

}
