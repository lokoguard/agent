package resource_monitoring

import (
	"strings"

	"github.com/shirou/gopsutil/host"
)

func FetchTempStats() ([]SensorTemperatureInfo, error) {
	sensors, err := host.SensorsTemperatures()

	if err != nil && !strings.Contains(err.Error(), "Number of warnings:") {
		return nil, err
	}

	tempInfo := make([]SensorTemperatureInfo, 0)

	for _, sensor := range sensors {
		if strings.Contains(sensor.SensorKey, "input") && sensor.Temperature != 0 {
			tempLabel := sensor.SensorKey
			label := strings.TrimSuffix(sensor.SensorKey, "_input")
			label = strings.TrimSuffix(label, "_thermal")
			if tempLabel != label {
				tempInfo = append(tempInfo, SensorTemperatureInfo{
					Sensor:             label,
					TemperatureCelcius: float32(sensor.Temperature),
				})
			}
		}
	}
	return tempInfo, nil
}
