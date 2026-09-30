package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"request-recorder/model"
)

func RequestHandler(res http.ResponseWriter, req *http.Request) {

	bodyData, err := io.ReadAll(req.Body)
	if err != nil {
		bodyData = []byte(fmt.Sprintf("error reading body: %v", err))
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	method := req.Method
	proto := req.Proto
	path := req.URL.Path
	query := req.URL.RawQuery
	headers := req.Header
	host := req.Host
	remoteAddr := req.RemoteAddr
	contentLength := req.ContentLength
	cookies := req.Header.Get("cookie")
	bodyStr := string(bodyData)

	logEntry := model.Logs{
		Timestamp:     timestamp,
		Method:        method,
		Proto:         proto,
		Path:          path,
		Query:         query,
		Host:          host,
		RemoteAddress: remoteAddr,
		Headers:       headers,
		ContentLength: contentLength,
		Cookies:       cookies,
		Body:          bodyStr,
	}

	res.WriteHeader(http.StatusOK)

	select {
	case model.LogQueue <- logEntry:
	default:
		fmt.Println("Warning, Log queue full! Dropping log to.stdout.")
		logEntryJson, jsonErr := json.Marshal(logEntry)
		if jsonErr != nil {
			fmt.Printf("error marshaling json log entry: %v", err)
		}
		fmt.Println(string(logEntryJson))
	}

}
