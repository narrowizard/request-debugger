package main

import "time"

type RequestData struct {
	URL         string
	Method      string
	Proto       string
	Header      map[string][]string
	Body        []byte
	CreatedTime time.Time
}
