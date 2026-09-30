package model

type Logs struct {
	Timestamp     string              `json:"timestamp"`
	Method        string              `json:"method"`
	Proto         string              `json:"proto"`
	Path          string              `json:"path"`
	Query         string              `json:"query"`
	Host          string              `json:"host"`
	RemoteAddress string              `json:"remoteAddress"`
	Headers       map[string][]string `json:"headers"`
	ContentLength int64               `json:"contentLength"`
	Cookies       string              `json:"cookies"`
	Body          string              `json:"body"`
}

var LogQueue = make(chan Logs, 10000)
