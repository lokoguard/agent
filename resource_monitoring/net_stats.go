package resource_monitoring

import (
	"errors"

	"github.com/shirou/gopsutil/net"
)

func FetchNetStats() (*NetStat, error) {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}
	IO := make(map[string][]int32)
	for _, IOStat := range netStats {
		nic := []int32{int32(IOStat.BytesSent), int32(IOStat.BytesRecv)}
		IO[IOStat.Name] = nic
	}
	if len(IO) == 0 {
		return &NetStat{
			BytesSent: 0,
			BytesRecv: 0,
		}, nil
	}
	if _, ok := IO["all"]; !ok {
		return nil, errors.New("eth0 not found")
	}
	allNet := IO["all"]
	currentBytesSent := allNet[0]
	currentBytesRecv := allNet[1]
	bytesSent := currentBytesSent - lastCurrentBytesSent
	bytesRecv := currentBytesRecv - lastCurrentBytesRecv
	lastCurrentBytesSent = currentBytesSent
	lastCurrentBytesRecv = currentBytesRecv
	if bytesSent <= 0 {
		bytesSent = 0
	}
	if bytesRecv <= 0 {
		bytesRecv = 0
	}
	return &NetStat{
		BytesSent: bytesSent,
		BytesRecv: bytesRecv,
	}, nil
}
