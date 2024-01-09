package main

import (
	"fmt"

	"github.com/lokoguard/agent/resource_monitoring"
)

// https://github.com/pesos/grofer/blob/main/pkg/metrics/general/serve_stats.go
// https://github.com/pesos/grofer/blob/main/pkg/metrics/general/general_stats.go

func main() {
	resourceStats, err := resource_monitoring.FetchResourceStats()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(resourceStats)
	fmt.Println(resourceStats.JSON())
}
