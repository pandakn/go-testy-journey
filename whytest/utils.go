package whytest

import "encoding/json"

func JsonToString(obj any) string {
	result, _ := json.Marshal(&obj)
	return string(result)
}
