package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"time"

	"../domain"
)

func updater(rate time.Duration) {
	for {
		req := domain.Request{
			Value: rand.Float64() * 100,
		}
		byt := req.JSON()
		buffer := bytes.NewBuffer(byt)
		http.Post("http://localhost:8080", "application/json", buffer)

		time.Sleep(rate)
	}
}

func main() {
	// inital delay
	time.Sleep(2 * time.Second)
	// start
	updater(1 * time.Second)
}
