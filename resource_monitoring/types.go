package resource_monitoring

import "fmt"

type ResourceStats struct {
	CPUStats  []float32               `json:"cpu_stats"`
	MemStats  MemoryStat              `json:"mem_stats"`
	DiskStats []DiskStat              `json:"disk_stats"`
	TempStats []SensorTemperatureInfo `json:"temp_stats"`
	NetStats  NetStat                 `json:"net_stats"`
	HostInfo  HostInfo                `json:"host_info"`
	TimeStamp uint64                  `json:"timestamp"`
}

func (r ResourceStats) String() string {
	cpuStatsString := ""
	for index, cpuStat := range r.CPUStats {
		cpuStatsString += fmt.Sprintf("CPU%d > %f\n", index+1, cpuStat)
	}
	diskStatString := ""
	for _, diskStat := range r.DiskStats {
		diskStatString += fmt.Sprintf("%s\n", diskStat)
	}
	tempStatString := ""
	for _, tempStat := range r.TempStats {
		tempStatString += fmt.Sprintf("%s\n", tempStat)
	}
	msg := "----- Resource Stats -----\n\n"
	msg += fmt.Sprintf("Host Info: \n%v\n\n", r.HostInfo)
	msg += fmt.Sprintf("CPU Stats: \n%s\n\n", cpuStatsString)
	msg += fmt.Sprintf("Memory Stats: \n%v\n\n", r.MemStats)
	msg += fmt.Sprintf("Disk Stats: \n%s\n\n", diskStatString)
	msg += fmt.Sprintf("Temperature Stats: \n%s\n\n", tempStatString)
	msg += fmt.Sprintf("Network Stats: \n%v\n\n", r.NetStats)
	msg += fmt.Sprintf("Timestamp: %d\n", r.TimeStamp)
	return msg
}

type HostInfo struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`    // seconds
	BootTime             uint64 `json:"boot_time"` // unix timestamp
	OS                   string `json:"os"`
	Platform             string `json:"platform"`
	PlatformFamily       string `json:"platform_family"`
	PlatformVersion      string `json:"platform_version"`
	KernelVersion        string `json:"kernel_version"`
	KernelArch           string `json:"kernel_arch"`
	VirtualizationSystem string `json:"virtualization_system"`
	VirtualizationRole   string `json:"virtualization_role"`
}

func (h HostInfo) String() string {
	return fmt.Sprintf("Hostname: %s\nUptime: %d\nBootTime: %d\nOS: %s\nPlatform: %s\nPlatformFamily: %s\nPlatformVersion: %s\nKernelVersion: %s\nKernelArch: %s\nVirtualizationSystem: %s\nVirtualizationRole: %s", h.Hostname, h.Uptime, h.BootTime, h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion, h.KernelVersion, h.KernelArch, h.VirtualizationSystem, h.VirtualizationRole)
}

type DiskStat struct {
	Path        string  `json:"path"`
	TotalGB     float32 `json:"total_gb"`
	UsedGB      float32 `json:"used_gb"`
	FreeGB      float32 `json:"free_gb"`
	UsedPercent float32 `json:"used_percent"`
	FSType      string  `json:"fs_type"`
}

func (d DiskStat) String() string {
	return fmt.Sprintf("Path: %s, Total: %f, Used: %f, Free: %f, UsedPercent: %f, FSType: %s", d.Path, d.TotalGB, d.UsedGB, d.FreeGB, d.UsedPercent, d.FSType)
}

type MemoryStat struct {
	TotalGB     float32 `json:"total_gb"`
	UsedGB      float32 `json:"used_gb"`
	AvailableGB float32 `json:"available_gb"`
	FreeGB      float32 `json:"free_gb"`
	CachedGB    float32 `json:"cached_gb"`
}

func (m MemoryStat) String() string {
	return fmt.Sprintf("Total: %f, Used: %f, Available: %f, Free: %f, Cached: %f", m.TotalGB, m.UsedGB, m.AvailableGB, m.FreeGB, m.CachedGB)
}

type SensorTemperatureInfo struct {
	Sensor             string  `json:"sensor"`
	TemperatureCelcius float32 `json:"temperature_celcius"`
}

func (s SensorTemperatureInfo) String() string {
	return fmt.Sprintf("Sensor: %s, Temperature: %f", s.Sensor, s.TemperatureCelcius)
}

type NetStat struct {
	BytesSent int32 `json:"bytes_sent"`
	BytesRecv int32 `json:"bytes_recv"`
}

func (n NetStat) String() string {
	return fmt.Sprintf("BytesSent: %d, BytesRecv: %d", n.BytesSent, n.BytesRecv)
}
