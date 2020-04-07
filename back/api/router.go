package api

import (
  "encoding/json"
  "net/http"
)


func AddRoute(url string, controller Controller) {
	var c = func(w http.ResponseWriter, req *http.Request){
		context := Context{Request: *req}

		data, _ := controller(context)
		jsonData, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
	http.HandleFunc(url, c)
}

