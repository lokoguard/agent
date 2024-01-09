package resource_monitoring

import "github.com/shirou/gopsutil/mem"

func FetchMemoryStats() (*MemoryStat, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return &MemoryStat{
		TotalGB:     float32(memory.Total) / 1024 / 1024 / 1024,
		UsedGB:      float32(memory.Used) / 1024 / 1024 / 1024,
		AvailableGB: float32(memory.Available) / 1024 / 1024 / 1024,
		FreeGB:      float32(memory.Free) / 1024 / 1024 / 1024,
		CachedGB:    float32(memory.Cached) / 1024 / 1024 / 1024,
	}, nil
}
