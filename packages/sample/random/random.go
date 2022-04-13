package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// upperBound is the upper bound of the random int range
	upperBound = 101
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Request is the function's request struct
type Request struct {
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	return &Response{
		Body: fmt.Sprintf("Hello, your random number is %d", rand.Intn(upperBound)),
	}, nil
}
