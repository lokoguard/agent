package resource_monitoring

import (
	"context"

	"github.com/shirou/gopsutil/host"
)

func FetchHostInfo() (*HostInfo, error) {
	info, err := host.InfoWithContext(context.Background())
	if err != nil {
		return nil, err
	}
	return &HostInfo{
		Hostname: info.Hostname,
		Uptime:   info.Uptime,
		BootTime: info.BootTime,
		OS:       info.OS,
		Platform: info.Platform,
		PlatformFamily: info.PlatformFamily,
		PlatformVersion: info.PlatformVersion,
		KernelVersion: info.KernelVersion,
		KernelArch: info.KernelArch,
		VirtualizationSystem: info.VirtualizationSystem,
		VirtualizationRole: info.VirtualizationRole,
	}, nil
}
