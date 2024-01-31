package resource_monitoring

import "encoding/json"

func (r ResourceStats) JSON() ([]byte, error) {
	if r.CPUStats == nil {
		r.CPUStats = []float32{}
	}
	if r.DiskStats == nil {
		r.DiskStats = []DiskStat{}
	}
	if r.TempStats == nil {
		r.TempStats = []SensorTemperatureInfo{}
	}
	jsonBytes , err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}