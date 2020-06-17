package api

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/url"

	env "github.com/devmarka/bbb-go-server/env"
)

func BBBCreateEvent(eventid string, logouturl string) string {
	event, _ := GetEvent(eventid)

	record := "false"
	if event.Record {
		record = "true"
	}

	arrParams := map[string]string{
		"name":            url.QueryEscape(event.EventName),
		"meetingID":       url.QueryEscape(event.Id),
		"attendeePW":      url.QueryEscape(event.GetAttendeePW()),
		"moderatorPW":     url.QueryEscape(event.GetModeratorPW()),
		"logoutURL":       url.QueryEscape(logouturl),
		"record":          record,
		"meta_suuid":      "",
		"welcome":         "",
		"maxParticipants": "30",
		"isBreakout":      "false",
	}

	url := BBBBuildURL("create", arrParams)
	response := HTTPResponse(url)
	if "ERROR" == response {
		log.Println("ERROR : HTTP ERROR")
		return "ERROR"
	}
	var data CreateMeetingResponse
	err := ResponseXML(response, &data)
	if err != nil {
		log.Println("XML ERROR")
		return "ERROR"
	}
	if "SUCCESS" == data.Returncode {
		return data.MeetingID
	}
	return "ERROR"
}

func BBBJoinMeetingURL(eventid string, username string, password string, logouturl string) (string, bool) {
	event, _ := GetEvent(eventid)
	if !event.Active {
		return "", false
	}
	if !BBBEventRunning(eventid) {
		_ = BBBCreateEvent(eventid, logouturl)
	}
	arrParams := map[string]string{
		"meetingID": url.QueryEscape(event.Id),
		"fullName":  url.QueryEscape(username),
		"password":  url.QueryEscape(password),
	}
	url := BBBBuildURL("join", arrParams)
	return url, true
}

func BBBEndEvent(eventid string) {

}

// BBBEventRunning::tested
func BBBEventRunning(eventid string) bool {
	arrParams := map[string]string{
		"meetingID": url.QueryEscape(eventid),
	}
	url := BBBBuildURL("isMeetingRunning", arrParams)
	response := HTTPResponse(url)
	if "ERROR" == response {
		log.Println("ERROR : HTTP ERROR")
		return false
	}
	var data IsMeetingRunningResponse
	err := ResponseXML(response, &data)
	if err != nil {
		log.Println("XML ERROR")
		return false
	}
	return data.Running
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
