package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"os"
	"time"

	"../domain"
	"github.com/codeship/go-retro"
	"github.com/pkg/errors"
)

var (
	url = os.Getenv("API_URL")
)

func updater(rate time.Duration) {
	for {
		req := domain.Request{
			Value: rand.Float64() * 100,
		}
		byt := req.JSON()
		buffer := bytes.NewBuffer(byt)

		finalErr := retro.DoWithRetry(func() error {
			_, err := http.Post(url, "application/json", buffer)
			return err
		})

		if finalErr != nil {
			panic(errors.Wrapf(finalErr, "Could not post to %v", url))
		}

		time.Sleep(rate)
	}
}

func main() {
	// inital delay
	time.Sleep(2 * time.Second)
	// start
	updater(1 * time.Second)
}
