package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"request-recorder/model"

	"request-recorder/api"
)

func init() {
	go func() {
		file, err := os.OpenFile("hits.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer file.Close()

		for logEntry := range model.LogQueue {
			logEntryJson, jsonErr := json.Marshal(logEntry)
			if jsonErr == nil {
				fmt.Println(string(logEntryJson))
				file.WriteString(string(logEntryJson) + "\n")
			}
		}
	}()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.RequestHandler)

	fmt.Println("Server running on localhost:8888")
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}

}
