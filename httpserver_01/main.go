package main

import (
	"GeekTrainingCamp/httpserver_01/controller"
	"encoding/json"
	"flag"
	glog "github.com/golang/glog"
	"net/http"
)

var MyRoute map[string]string = make(map[string]string, 3)

func main() {
	portParam := flag.String("p", "8080", "Server Listen Port")
	flag.Parse()
	defer glog.Flush()
	registerRoutes()
	serverAddr := ":" + *portParam
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		glog.Error("Start server error:", err)
	}
}

func serverHandle(response http.ResponseWriter, request *http.Request) {
	var str []string
	for _, path := range MyRoute {
		str = append(str, path)
	}
	result, _ := json.Marshal(str)
	response.Write(result)
}

func registerRoutes() {
	registerRoute("/", handleIterceptor(serverHandle), nil)
	registerRoute("/healthz", handleIterceptor(controller.Healthz), nil)
	registerRoute("/filter/getHeader", handleIterceptor(controller.GetRequestHeaderHandle), nil)
	registerRoute("/filter/getSystemGoVersion", handleIterceptor(controller.GetSystemGoVersion), nil)
	versionParam := []string{"sysConfigName"}
	registerRoute("/filter/getSystemEnv", handleIterceptor(controller.GetSystemEnv), versionParam)
}

func registerRoute(path string, handler func(http.ResponseWriter, *http.Request), headParams []string) {
	var pathStr string
	if headParams != nil {
		pathStr = path + "?"
		for _, param := range headParams {
			pathStr += param + "= &"
		}
		if pathStr[len(pathStr)-1] == '&' {
			pathStr = pathStr[0 : len(pathStr)-1]
		}
	} else {
		pathStr = path
	}
	MyRoute[path] = pathStr
	http.HandleFunc(path, handler)
}

func handleIterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
		glog.Info("Client IP:", r.RemoteAddr)
		glog.Info("Request Path:", r.RequestURI)
		glog.Info("Response HTTP Code:", w.Header().Get("Status"))
	}
}
