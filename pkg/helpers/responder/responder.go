package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func Success(w http.ResponseWriter, data interface{}, message string) {
	var response Response
	response.Code = 2200
	response.Data = data
	response.Message = message

	re, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(re)
	return
}

func Error(w http.ResponseWriter, Code int, error interface{}) {
	var response Response
	response.Code = Code
	response.Data = []string{}
	response.Error = error

	re, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Code)
	w.Write(re)
	return
}
