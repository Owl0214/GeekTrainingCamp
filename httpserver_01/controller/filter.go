package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

type Header struct {
}

func GetRequestHeaderHandle(response http.ResponseWriter, request *http.Request) {
	requestHeader := request.Header
	result, _ := json.Marshal(requestHeader)
	DoResponse(response, result)
}

func GetSystemGoVersion(response http.ResponseWriter, request *http.Request) {
	version := runtime.Version()
	result, _ := json.Marshal(version)
	DoResponse(response, result)
}

func GetSystemEnv(response http.ResponseWriter, request *http.Request) {
	var findResult string
	values := request.URL.Query()
	configName := values.Get("sysConfigName")
	fmt.Println(configName)
	env := os.Getenv(configName)
	if env != "" {
		findResult = env
	} else {
		findResult = configName + "not found"
	}
	result, _ := json.Marshal(findResult)
	DoResponse(response, result)
}

func DoResponse(response http.ResponseWriter, result []byte) {
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Status", fmt.Sprintf("%d", http.StatusOK))
	response.Write(result)
}
