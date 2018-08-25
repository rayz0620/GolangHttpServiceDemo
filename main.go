package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HTTP_HEADER_CONTENT_TYPE = "content-type"
	CONTENT_TYPE_JSON        = "application/json"
	DEFAULT_PORT             = ":8080"
)

func readFileHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	timeout := request.PostForm.Get("timeout")
	path := request.PostForm.Get("path")

	timeBeforeRead := time.Now()
	ioutil.ReadFile(path)
	timeAfterRead := time.Now()
	duration := timeAfterRead.Sub(timeBeforeRead)

	serverTime := time.Now()
	respDict := map[string]string{
		"time": serverTime.Format("2006-01-02 15:04:05"),
		"timeout":    timeout,
		"path":       path,
		"duration":   fmt.Sprintf("%d", duration.Nanoseconds()/1000),
	}

	respBody, err := json.Marshal(respDict)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add(HTTP_HEADER_CONTENT_TYPE, CONTENT_TYPE_JSON)
	writer.Write(respBody)
}

func main() {
	fmt.Printf("Http serving at http://127.0.0.1%s\n", DEFAULT_PORT)

	http.HandleFunc("/", readFileHandler)
	http.ListenAndServe(DEFAULT_PORT, nil)
}
