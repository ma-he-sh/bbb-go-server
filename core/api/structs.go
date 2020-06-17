package api

// code ref from :: https://github.com/ypgao1/bigbluebutton-api-go

// GetMeetingJoinResponse::
type GetMeetingJoinResponse struct {
	ReturnCode   string `xml:"returncode"`
	MessageKey   string `xml:"messageKey"`
	Message      string `xml:"message"`
	MeetingID    string `xml:"meeting_id"`
	UserID       string `xml:"user_id"`
	AuthToken    string `xml:"auth_token"`
	SessionToken string `xml:"session_token"`
	URL          string `xml:"url"`
}

// IsMeetingRunningResponse::
type IsMeetingRunningResponse struct {
	ReturnCode string `xml:"returncode"`
	Running    bool   `xml:"running"`
}

// CreateMeetingResponse::
type CreateMeetingResponse struct {
	Returncode           string `xml:"returncode"`
	MeetingID            string `xml:"meetingID"`
	CreateTime           string `xml:"createTime"`
	AttendeePW           string `xml:"attendeePW"`
	ModeratorPW          string `xml:"moderatorPW"`
	HasBeenForciblyEnded string `xml:"hasBeenForciblyEnded"`
	MessageKey           string `xml:"messageKey"`
	Message              string `xml:"message"`
}

type EndResponse struct {
	ReturnCode string `xml:"returncode"`
	MessageKey string `xml:"messageKey"`
	Message    string `xml:"message"`
}
type GetMeetingsResponse struct {
	ReturnCode string      `xml:"returncode"`
	Meetings   allMeetings `xml:"meetings"`
}
type allMeetings struct {
	MeetingInfo []GetMeetingInfoResponse `xml:"meeting"`
}

type GetMeetingInfoResponse struct {
	ReturnCode            string    `xml:"returncode"`
	MeetingName           string    `xml:"meetingName"`
	MeetingID             string    `xml:"meetingID"`
	InternalMeetingID     string    `xml:"internalMeetingID"`
	CreateTime            string    `xml:"createTime"`
	CreateDate            string    `xml:"createDate"`
	VoiceBridge           string    `xml:"voiceBridge"`
	DialNumber            string    `xml:"dialNumber"`
	AttendeePW            string    `xml:"attendeePW"`
	ModeratorPW           string    `xml:"moderatorPW"`
	Running               bool      `xml:"running"`
	Duration              int       `xml:"duration"`
	HasUserJoined         bool      `xml:"hasUserJoined"`
	Recording             bool      `xml:"recording"`
	HasBeenForciblyEnded  bool      `xml:"hasBeenForciblyEnded"`
	StartTime             string    `xml:"startTime"`
	EndTime               string    `xml:"endTime"`
	ParticipantCount      int       `xml:"participantCount"`
	ListenerCount         int       `xml:"listenerCount"`
	VoiceParticipantCount int       `xml:"voiceParticipantCount"`
	VideoCount            int       `xml:"videoCount"`
	MaxUsers              int       `xml:"maxUsers"`
	ModeratorCount        int       `xml:"moderatorCount"`
	Attendees             attendees `xml:"attendees"`
	Metadata              string    `xml:"metadata"`
	MessageKey            string    `xml:"messageKey"`
	Message               string    `xml:"message"`
	//untested
	BreakoutRooms breakoutRooms `xml:"breakoutRooms"`
}
type breakoutRooms struct {
	BreakoutRooms []string `xml:"breakout"`
}

type attendees struct {
	Attendees []attendee `xml:"attendee"`
}

type attendee struct {
	UserID          string `xml:"userID"`
	FullName        string `xml:"fullName"`
	Role            string `xml:"role"`
	IsPresenter     bool   `xml:"isPresenter"`
	IsListeningOnly bool   `xml:"isListeningOnly"`
	HasJoinedVoice  bool   `xml:"hasJoinedVoice"`
	HasVideo        bool   `xml:"hasVideo"`
	Customdata      string `xml:"customdata"`
}

type GetRecordingsResponse struct {
	ReturnCode string     `xml:"returncode"`
	Recordings recordings `xml:"recordings"`
}
type recordings struct {
	Recording []recording `xml:"recording"`
}
type recording struct {
	RecordID     string   `xml:"recordID"`
	MeetingID    string   `xml:"meetingID"`
	Name         string   `xml:"name"`
	Published    string   `xml:"published"`
	State        string   `xml:"state"`
	StartTime    string   `xml:"startTime"`
	EndTime      string   `xml:"endTime"`
	Participants string   `xml:"participants"`
	MetaData     metadata `xml:"metadata"`
	Playback     struct {
		Format []struct {
			Type   string   `xml:"type"`
			Url    string   `xml:"url"`
			Length string   `xml:"length"`
			Images []string `xml:"preview>images>image"`
		} `xml:"format"`
	} `xml:"playback"`
}

type metadata struct {
	Title       string `xml:"title"`
	Subject     string `xml:"subject"`
	Description string `xml:"description"`
	Creator     string `xml:"creator"`
	Contributor string `xml:"contributor"`
	Language    string `xml:"language"`
}

type CreateWebhookResponse struct {
	Returncode string `xml:"returncode"`
	MessageKey string `xml:"messageKey"`
	Message    string `xml:"message"`
	HookID     string `xml:"hookID"`
}
type DestroyedWebhookResponse struct {
	Returncode string `xml:"returncode"`
	MessageKey string `xml:"messageKey"`
	Message    string `xml:"message"`
	Removed    string `xml:"removed"`
}

type PublishRecordingsResponse struct {
	ReturnCode string `xml:"returncode"`
	Published  string `xml:"published"`
}

type DeleteRecordingsResponse struct {
	ReturnCode string `xml:"returncode"`
	Deleted    string `xml:"deleted"`
}

type Recording struct {
	MeetingID string
	RecordID  string
	State     string
	Meta      string
	Publish   string
}

type Participants struct {
	IsAdmin_     int
	FullName_    string
	MeetingID_   string
	Password_    string
	CreateTime   string
	UserID       string
	WebVoiceConf string
	ConfigToken  string
	AvatarURL    string
	Redirect     string
	ClientURL    string
	JoinURL      string
}

type MeetingRoom struct {
	Name_                       string
	MeetingID_                  string
	InternalMeetingId           string
	AttendeePW_                 string
	ModeratorPW_                string
	Welcome                     string
	DialNumber                  string
	VoiceBridge                 string
	WebVoice                    string
	LogoutURL                   string
	Record                      string
	Duration                    int
	Meta                        string
	ModeratorOnlyMessage        string
	AutoStartRecording          bool
	AllowStartStopRecording     bool
	Created                     bool
	PostId                      string
	CreatedAt                   int64
	EndedAt                     int64
	AttendeeNames               []string
	LoopCount                   int
	Meta_bn_recording_ready_url string
	Meta_channelid              string
	Meta_endcallbackurl         string
	CreateMeetingResponse       CreateMeetingResponse
	MeetingInfo                 GetMeetingInfoResponse
}

type WebHook struct {
	HookID          string
	CallBackURL     string
	MeetingId       string
	WebhookResponse CreateWebhookResponse
}
