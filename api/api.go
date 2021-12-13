package api

import (
	"assignment/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func GetMostUsedWords(writer http.ResponseWriter, request *http.Request) {

	var req Request
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		logrus.Error("invalid input")
		utils.ErrorResponse(writer, "invalid input", 400)
		return
	}

	reader := strings.NewReader(req.Text)
	words := countWords(reader)

	utils.SuccessResponse(writer, words, 200)
}
