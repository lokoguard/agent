package resource_monitoring

import "time"

func FetchResourceStats() (*ResourceStats, error) {
	resourceStats := &ResourceStats{}
	cpuStats, err := FetchCPUStats()
	if err != nil {
		return nil, err
	}
	resourceStats.CPUStats = cpuStats
	memStats, err := FetchMemoryStats()
	if err != nil {
		return nil, err
	}
	resourceStats.MemStats = *memStats
	diskStats, err := FetchDiskStats()
	if err != nil {
		return nil, err
	}
	resourceStats.DiskStats = diskStats
	tempStats, err := FetchTempStats()
	if err != nil {
		return nil, err
	}
	resourceStats.TempStats = tempStats
	netStats, err := FetchNetStats()
	if err != nil {
		return nil, err
	}
	resourceStats.NetStats = *netStats
	hostInfo, err := FetchHostInfo()
	if err != nil {
		return nil, err
	}
	resourceStats.HostInfo = *hostInfo
	resourceStats.TimeStamp = uint64(time.Now().Unix())
	return resourceStats, nil
}
