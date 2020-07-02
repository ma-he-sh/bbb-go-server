package templ

import (
	"math/rand"
)

const bucket = "ABCDEFGHIJKLMNOPQRSTUVW123456789"
const bucketStrLen = 8

func headerSection() string {
	return `<div class='nav--wrapper'>
		<a href='/admin/dashboard' class='nav--item' data-tab='admin-tab-event'>Events</a>
		<a href='/admin/dashboard?tab=recordings' class='nav--item' data-tab='admin-tab-recordings'>Recordings</a>
		<a href='/admin/dashboard?tab=reports' class='nav--item' data-tab='admin-tab-reports'>Reports</a>
	</div>
	<div class='nav--buttons'>
		<a href='/admin/signout' class='nav--button'>
			<svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-log-out'><path d='M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4'></path><polyline points='16 17 21 12 16 7'></polyline><line x1='21' y1='12' x2='9' y2='12'></line></svg>
		</a>
		<button class='nav--button'>
			<svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-settings'><circle cx='12' cy='12' r='3'></circle><path d='M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z'></path></svg>	
		</button>
	</div>
	`
}

func eventTemplate(data map[string]interface{}) string {
	return `<div class='bk--content' data-content='admin-tab-event'>
		<div class='bk--content-wrapper'>
			<form id='admin_create_event'> 	
				<input id='form_action' type='hidden' value='create_event'/>
				<div class='bk--cont-header'>
					Add Event
				</div>
				<div class='bk--cont-body'>
					` + InputToggle(ToggleStruct{Id: "toggle_active", Name: "toggle_active", Label: "Active", Checked: data["toggle_active"].(bool), Atts: nil}) +
		Input(InputStruct{Type: "text", Id: "str_event_name", Name: "str_event_name", Label: "Event Name", Value: data["str_event_name"].(string), Atts: nil}) +
		Input(InputStruct{Type: "text", Id: "str_event_message", Name: "str_event_message", Label: "Event Message", Value: data["str_event_message"].(string), Atts: nil}) + `
					<div class='input--group'>
						` + Input(InputStruct{Type: "text", Id: "str_moderator_pw", Name: "str_moderator_pw", Label: "Moderator Passw", Value: data["str_moderator_pw"].(string), Atts: nil}) +
		Input(InputStruct{Type: "text", Id: "str_attendee_pw", Name: "str_attendee_pw", Label: "Attendee Passw", Value: data["str_attendee_pw"].(string), Atts: nil}) + `
					</div>
					` + InputToggle(ToggleStruct{Id: "toggle_record", Name: "toggle_record", Label: "Record", Checked: data["toggle_record"].(bool), Atts: nil}) + `
				</div>
				<div class='bk--cont-footer'>
					<div class='button--wrapper' style='width:200px;'>
					` + BKButton(ButtonStruct{Type: "submit", Varient: "primary", Id: "btn_create_event", Name: "btn_create_event", Label: "Create Event", Active: true, Atts: nil}) + `
					</div>
				</div>
			</form>
		</div>
	</div>`
}

func eventTemplWrapper(i int, event map[string]interface{}) string {
	status := "Not Active"
	if event["eventActive"].(bool) {
		status = "Active"
	}

	eventLink := `/event/` + event["eventid"].(string) + `/`
	joinLink := `/admin/dashboard/join/` + event["eventid"].(string) + `/` 
	shareLink := event["domain"].(string) + eventLink

	card := `<div class='app--event-wrapper'>
		<div class='event--header'>
			<div class='event--title'>` + event["eventName"].(string) + `</div>
			<div class='event--status'>
				<div class='event--active'>` + status + `</div>
			</div>`

	if event["eventRecord"].(bool) {
		card += `<div class='event--status'>
					<div class='event--active'>Recording</div>
				</div>`
	}

	card += `</div>
		<div class='event--body'>
			<div class='event--info-section'>
				<div class='section--title'>Attendee PW:</div>
				<div class='section--value'><code>` + event["attendeePW"].(string) + `</code></div>
			</div>
			<div class='event--info-section'>
				<div class='section--title'>Moderator PW:</div>
				<div class='section--value'><code>` + event["moderatorPW"].(string) + `</code></div>
			</div>
			<div class='event--info-section'>
				<div class='input--wrapper'>
				`+Input(InputStruct{Type: "text", Id: "", Name: "", Label: "Share Link", Value: shareLink, Atts: nil}) +`
				</div>
			</div>
			<div class='button--wrapper' style='width:200px;margin-right:10px;'>
				<a target='_self' href='` + joinLink + `' class='btn--mini'>Join</a>
			</div>
			<div class='button--wrapper' style='width:200px;'>
				<a target='_blank' href='` + eventLink + `' class='btn--mini'>Preview</a>
			</div>
		</div>
	</div>`

	return card
}

func eventList(events []map[string]interface{}) string {
	eventTempl := `<div class='bk--content' data-content='admin-tab-recordings'>
		<div class='bk--content-wrapper'>
			<div class='bk--cont-header'>
				<div class='button--wrapper' style='width:200px;'>
				` + BKLinkButton(LinkStruct{Varient: "primary", Name: "btn_create_event", Id: "btn_create_event", Label: "Create Event", Link: "/admin/dashboard/event/add", Atts: nil}) + `
				</div>
			</div>
			<div class='bk--cont-body'>`

	for i, event := range events {
		eventTempl += eventTemplWrapper(i, event)
	}

	eventTempl += ""

	eventTempl += `</div>
		</div>
	</div>`
	return eventTempl
}

type PagePayload struct {
	Page      string
	EventList []map[string]interface{}
}

func AdminDashboard(payload PagePayload) string {
	page := `<div class='app--backend'>
			<div class='bk--header'>` + headerSection() + `</div>
			<div class='bk--body'>`

	switch payload.Page {
	case "create_event":
		page += newEvent()
		break
	case "recordings":
		page += ""
		break
	case "reports":
		page += ""
		break
	default:
		page += eventList(payload.EventList)
	}
	page += `</div>`

	return backend(page)
}

func createToken() string {
	b := make([]byte, bucketStrLen)
	for i := range b {
		b[i] = bucket[rand.Intn(len(bucket))]
	}
	return string(b)
}

func newEvent() string {
	moderatorPW := createToken()
	attendeePW := createToken()
	if moderatorPW == attendeePW {
		attendeePW = createToken()
	}

	eventData := map[string]interface{}{
		"toggle_active":     true,
		"str_event_name":    "",
		"str_moderator_pw":  moderatorPW,
		"str_attendee_pw":   attendeePW,
		"str_event_message": "",
		"toggle_record":     false,
	}

	return eventTemplate(eventData)
}

func editEvent() string {
	eventData := map[string]interface{}{
		"toggle_active":     true,
		"str_event_name":    "",
		"str_moderator_pw":  "",
		"str_attendee_pw":   "",
		"str_event_message": "",
		"toggle_record":     false,
	}

	return eventTemplate(eventData)
}
