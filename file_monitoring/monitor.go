package file_monitoring

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"k8s.io/utils/inotify"
)

func NewMonitor(callback ResultCallbackType) *Monitor {
	return &Monitor{
		CancelFuncs:    map[string]context.CancelFunc{},
		ResultCallback: callback,
		Mutex:          sync.RWMutex{},
	}
}

func (m *Monitor) AddPath(fileName string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if m.CancelFuncs[fileName] != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	m.CancelFuncs[fileName] = cancel
	go m.startMonitoring(ctx, fileName)
}

func (m *Monitor) RemovePath(fileName string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if m.CancelFuncs[fileName] == nil {
		return
	}
	m.CancelFuncs[fileName]()
	delete(m.CancelFuncs, fileName)
}

func (m *Monitor) SetFileList(fileList []string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	// find out which files to remove
	for fileName, cancel := range m.CancelFuncs {
		if !contains(fileList, fileName) {
			cancel()
			delete(m.CancelFuncs, fileName)
		}
	}
	// find out which files to add
	for _, fileName := range fileList {
		if m.CancelFuncs[fileName] == nil {
			ctx, cancel := context.WithCancel(context.Background())
			m.CancelFuncs[fileName] = cancel
			go m.startMonitoring(ctx, fileName)
		}
	}
}

func (m *Monitor) startMonitoring(ctx context.Context, fileName string) {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			if !(strings.HasSuffix(ev.Name, ".cache") ||
				strings.HasSuffix(ev.Name, ".swp") ||
				strings.HasSuffix(ev.Name, ".swx") ||
				strings.HasSuffix(ev.Name, ".swpx") ||
				strings.HasSuffix(ev.Name, ".alias")) && (ev.Mask&inotify.InCreate == inotify.InCreate ||
				ev.Mask&inotify.InAccess == inotify.InAccess ||
				ev.Mask&inotify.InDelete == inotify.InDelete ||
				ev.Mask&inotify.InCloseWrite == inotify.InCloseWrite) {
				var eventType EventType
				switch {
				case ev.Mask&inotify.InCreate == inotify.InCreate:
					eventType = Create
				case ev.Mask&inotify.InAccess == inotify.InAccess:
					eventType = Access
				case ev.Mask&inotify.InDelete == inotify.InDelete:
					eventType = Delete
				case ev.Mask&inotify.InCloseWrite == inotify.InCloseWrite:
					eventType = Write
				}
				m.ResultCallback(&UpdateEvent{
					FileName:  ev.Name,
					Type:      eventType,
					Timestamp: time.Now().Unix(),
				})
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
			return
		case <-ctx.Done():
			return
		}
	}
}


func contains(fileList []string, fileName string) bool {
	for _, f := range fileList {
		if f == fileName {
			return true
		}
	}
	return false
}