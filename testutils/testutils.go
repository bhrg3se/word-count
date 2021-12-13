package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetPOSTRequest(data interface{}) *http.Request {
	mar, _ := json.Marshal(data)
	buff := bytes.NewBuffer(mar)
	r, _ := http.NewRequest("POST", "/", buff)
	return r
}
