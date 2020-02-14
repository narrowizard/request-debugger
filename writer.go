package main

import "fmt"

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	NewLine      = "\n"
)

type Writer struct {
	DataList chan RequestData
}

func NewWriter() *Writer {
	var w = &Writer{
		DataList: make(chan RequestData, 100),
	}
	go func() {
		for {
			if data, ok := <-w.DataList; ok {
				// start debug
				fmt.Printf(ErrorColor+" at %s"+NewLine, "receive request", data.CreatedTime.String())
				// url
				fmt.Printf(InfoColor+NewLine, "-----url-----")
				fmt.Println(data.URL)
				// method
				fmt.Printf(InfoColor+NewLine, "----method----")
				fmt.Println(data.Method)
				// proto
				fmt.Printf(InfoColor+NewLine, "----proto----")
				fmt.Println(data.Proto)
				// headers
				fmt.Printf(InfoColor+NewLine, "---headers---")
				for k, v := range data.Header {
					fmt.Printf("%s: %s\n", k, v)
				}
				// body
				fmt.Printf(InfoColor+NewLine, "----body----")
				fmt.Println(string(data.Body))
				fmt.Println("")
			}
		}
	}()
	return w
}

func (w *Writer) PushData(d RequestData) {
	w.DataList <- d
}
