package syslogserver


type ResultCallbackType func (*SyslogMessage, error)

type SyslogMessage struct {
	Version         uint16
	FacilityMessage string
	FacilityLevel   string
	SeverityMessage string
	SeverityLevel   string
	Hostname        string
	Appname         string
	Message         string
	Timestamp	   int
}


