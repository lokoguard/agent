package file_monitor

import (
	"context"
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

type Monitor struct {
	CancelFuncs    map[string]context.CancelFunc
	ResultCallback ResultCallbackType
	Mutex          sync.RWMutex
}

type ResultCallbackType func(event *UpdateEvent)
