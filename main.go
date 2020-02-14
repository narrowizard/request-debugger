package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var w = NewWriter()

func main() {
	var port = ":8080"
	http.HandleFunc("/", handler)
	fmt.Println("start listening at " + port)
	http.ListenAndServe(port, nil)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	var m RequestData
	m.CreatedTime = time.Now()
	m.URL = req.URL.String()
	m.Method = req.Method
	m.Proto = req.Proto
	m.Header = req.Header
	var body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		w.PushData(m)
		fmt.Printf(ErrorColor+NewLine, err)
		writer.WriteHeader(500)
		return
	}
	defer req.Body.Close()
	m.Body = body
	w.PushData(m)

	writer.WriteHeader(200)
}
