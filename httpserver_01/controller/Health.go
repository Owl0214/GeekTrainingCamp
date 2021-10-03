package controller

import (
	"encoding/json"
	"net/http"
)

func Healthz(response http.ResponseWriter, request *http.Request) {
	str := "{Code:200,Msg:'Serving'}"
	result, _ := json.Marshal(str)
	DoResponse(response, result)
}
