package server

import (
	"encoding/json"
)

func ParseAlert(data []byte) (Alert, error) {
	var alert Alert
	if err := json.Unmarshal(data, &alert); err != nil {
		return Alert{}, err
	}
	return alert, nil
}
