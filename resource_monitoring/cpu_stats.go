package resource_monitoring

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func FetchCPUStats() ([]float32, error) {
	cpuRates, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	cpuRatesFloat32 := make([]float32, len(cpuRates))
	for i, rate := range cpuRates {
		cpuRatesFloat32[i] = float32(rate)
	}
	return cpuRatesFloat32, nil
}
