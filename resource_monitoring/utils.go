package resource_monitoring

import "encoding/json"

func (r ResourceStats) JSON() (string, error) {
	jsonString , err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}