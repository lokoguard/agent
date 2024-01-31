package syslogserver

import (
	"fmt"
	"time"

	"github.com/influxdata/go-syslog/rfc5424"
)

func FormatSyslogMessage(msg string) (SyslogMessage, error) {
	msgBytes := []byte(msg)
	p := rfc5424.NewParser()
	m, err := p.Parse(msgBytes, nil)
	if err != nil {
		return SyslogMessage{}, err
	}
	if !m.Valid() {
		return SyslogMessage{}, fmt.Errorf("invalid syslog message")
	}
	return SyslogMessage{
		Version:         m.Version(),
		FacilityMessage: handleNilString(m.FacilityMessage()),
		FacilityLevel:   handleNilString(m.FacilityLevel()),
		SeverityMessage: handleNilString(m.SeverityMessage()),
		SeverityLevel:   handleNilString(m.SeverityLevel()),
		Hostname:        handleNilString(m.Hostname()),
		Appname:         handleNilString(m.Appname()),
		Message:         handleNilString(m.Message()),
		Timestamp:       handleNilTimestamp(m.Timestamp()), // seconds since epoch
	}, nil
}

func handleNilString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func handleNilTimestamp(t *time.Time) int {
	if t == nil {
		return 0
	}
	return int(t.Unix())
}
