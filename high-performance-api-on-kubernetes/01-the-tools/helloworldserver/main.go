package main

import (
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, res *http.Request) {
		id := uuid.NewV4().String()
		log.Printf("[%v] Request received", id)
		now := time.Now()
		resp.Write([]byte("Hello World!"))
		log.Printf("[%v] Request execution time: %v", id, time.Now().Sub(now))
	})

	log.Println("Starting helloworld server")
	http.ListenAndServe(":8080", nil)
}
