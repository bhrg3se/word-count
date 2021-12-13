package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	resp := ApiResponse{
		Status:  true,
		Message: "",
		Data:    data,
	}
	marshalledResp, _ := json.Marshal(resp)
	writer.Write(marshalledResp)
}

func ErrorResponse(writer http.ResponseWriter, message string, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := ApiResponse{Status: false, Message: message}
	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
