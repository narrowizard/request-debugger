package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	NewLine      = "\n"
)

func main() {
	var port = ":80"
	http.HandleFunc("/", handler)
	fmt.Println("start listening at " + port)
	http.ListenAndServe(port, nil)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	// start debug
	fmt.Printf(ErrorColor+NewLine, "receive request")
	// url
	fmt.Printf(InfoColor+NewLine, "-----url-----")
	fmt.Println(req.URL)
	// method
	fmt.Printf(InfoColor+NewLine, "----method----")
	fmt.Println(req.Method)
	// proto
	fmt.Printf(InfoColor+NewLine, "----proto----")
	fmt.Println(req.Proto)
	// headers
	fmt.Printf(InfoColor+NewLine, "---headers---")
	for k, v := range req.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	// body
	var body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf(ErrorColor+NewLine, err)
		writer.WriteHeader(500)
		return
	}
	defer req.Body.Close()
	fmt.Printf(InfoColor+NewLine, "----body----")
	fmt.Println(string(body))
	fmt.Println("")
	writer.WriteHeader(200)
}
