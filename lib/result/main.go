package result

import (
	"encoding/json"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json_return(success bool, message string, data interface{}) []byte {
	response := Response{
		Success: success,
		Message: message,
		Data:    data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {

		jsonData, _ = json.Marshal(Response{
			Success: success,
			Message: "Unable to parse result!",
			Data:    data,
		})
		return jsonData
	} else {
		return jsonData
	}
}
