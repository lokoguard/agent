package main

import (
	"log"
	"sync"

	"github.com/lokoguard/agent/script_executor"
)

// https://github.com/pesos/grofer/blob/main/pkg/metrics/general/serve_stats.go
// https://github.com/pesos/grofer/blob/main/pkg/metrics/general/general_stats.go

func main() {

	s := script_executor.ScriptDefinition{
		TaskID: "test",
		Script: "sudo ls /",
		Args:   []string{},
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	s.RunWithCallback(func(res *script_executor.ScriptResult, err error) {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
		wg.Done()
	})
	wg.Wait()
	// res, err := s.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(res)

	// watcher, err := inotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = watcher.Watch("/etc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	select {
	// 	case ev := <-watcher.Event:
	// 		log.Println("event:", ev)
	// 	case err := <-watcher.Error:
	// 		log.Println("error:", err)
	// 	}
	// }
}
