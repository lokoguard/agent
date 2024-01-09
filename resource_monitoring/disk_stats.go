package resource_monitoring

import (
	"strings"

	"github.com/shirou/gopsutil/disk"
)

func FetchDiskStats() ([]DiskStat, error) {
	var partitions []disk.PartitionStat
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var diskStats []DiskStat = make([]DiskStat, 0)
	for _, value := range partitions {
		if strings.HasPrefix(value.Device, "/dev/loop") {
			continue
		} else if strings.HasPrefix(value.Mountpoint, "/var/lib/docker") {
			continue
		}
		usageVals, err := disk.Usage(value.Mountpoint)
		if err != nil {
			continue
		}
		diskStats = append(diskStats, DiskStat{
			Path:        usageVals.Path,
			TotalGB:     float32(usageVals.Total) / (1024 * 1024 * 1024),
			UsedGB:      float32(usageVals.Used) / (1024 * 1024 * 1024),
			UsedPercent: float32(usageVals.UsedPercent),
			FreeGB:      float32(usageVals.Free) / (1024 * 1024 * 1024),
			FSType:      usageVals.Fstype,
		})
	}

	return diskStats, nil
}
