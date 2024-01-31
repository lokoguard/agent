package file_monitoring

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type EventType string

const (
	Create EventType = "create"
	Access EventType = "access"
	Delete EventType = "delete"
	Write  EventType = "write"
)

type UpdateEvent struct {
	FileName  string    `json:"file_name"`
	Type      EventType `json:"type"`
	Timestamp int64     `json:"timestamp"`
}

func (event UpdateEvent) String() string {
	return event.FileName + " " + string(event.Type) + " " + fmt.Sprint(event.Timestamp)
}

func (event UpdateEvent) JSON() ([]byte, error) {
	return json.Marshal(event)
}

type Monitor struct {
	CancelFuncs    map[string]context.CancelFunc
	ResultCallback ResultCallbackType
	Mutex          sync.RWMutex
}

type ResultCallbackType func(event *UpdateEvent)
