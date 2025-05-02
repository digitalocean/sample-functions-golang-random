package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	envAppUrl = "APP_URL"
)

var (
	hostname = ""
)

// Greeting wraps the greeting message from the serverless backend.
type Greeting struct {
	Greeting                    string `json:"greeting"`
	FunctionDurationMillisecond int64  `json:"function_duration_millisecond"`
}

func init() {
	hostname = os.Getenv(envAppUrl)
	if hostname == "" {
		log.Fatalf("Environment variable %q not found\n", envAppUrl)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		url := fmt.Sprintf("%s/functions/nodejs/say-it-sammy", hostname)

		res, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to send request to serverless backend"))
			return
		}

		if res.StatusCode != http.StatusOK {
			w.WriteHeader(res.StatusCode)
			io.Copy(w, res.Body)
			return
		}

		greeting := &Greeting{}
		if err := json.NewDecoder(res.Body).Decode(greeting); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to decode response from serverless backend"))
			return
		}

		if greeting.Greeting == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("received empty greeting from serverless backend"))
			return
		}

		durationMillis := time.Now().Sub(start).Milliseconds()
		greeting.FunctionDurationMillisecond = durationMillis
		if err := json.NewEncoder(w).Encode(greeting); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to encode json response"))
			return
		}
	})

	log.Printf("starting up service with hostname %q\n", hostname)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start service: %s", err.Error())
	}
}
