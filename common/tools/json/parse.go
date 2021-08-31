package json

import (
	"encoding/json"
	"log"
)

func ParseJson(jsonStr string) map[string]interface{} {
	var response map[string]interface{}
	if string(jsonStr) != "" {
		err := json.Unmarshal([]byte(string(jsonStr)), &response)
		if err != nil {
			log.Println("生成json字符串错误")
		}
	}
	return response
}
