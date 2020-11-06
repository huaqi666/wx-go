package util

import "encoding/json"

func ToMap(request interface{}) map[string]interface{} {
	b, err := json.Marshal(request)
	if err != nil {
		return nil
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil
	}
	return data
}
