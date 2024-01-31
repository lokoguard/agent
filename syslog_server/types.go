package syslogserver

type ResultCallbackType func(*SyslogMessage, error)

type SyslogMessage struct {
	Version         uint16 `json:"version"`
	FacilityMessage string `json:"facility_message"`
	FacilityLevel   string `json:"facility_level"`
	SeverityMessage string `json:"severity_message"`
	SeverityLevel   string `json:"severity_level"`
	Hostname        string `json:"hostname"`
	Appname         string `json:"appname"`
	Message         string `json:"message"`
	Timestamp       int    `json:"timestamp"`
}
