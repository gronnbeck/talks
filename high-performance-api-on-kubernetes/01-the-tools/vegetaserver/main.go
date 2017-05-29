package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"./rollingwindow"

	"github.com/pkg/errors"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	targetURL := os.Getenv("TARGET_URL")

	if targetURL == "" {
		panic("Must specify TARGET_URL")
	}

	window := rollingwindow.New(1000)

	go func() {
		for {
			rate := uint64(1000) // per second
			duration := 4 * time.Second
			targeter := vegeta.NewStaticTargeter(vegeta.Target{
				Method: "GET",
				URL:    targetURL,
			})
			attacker := vegeta.NewAttacker()

			for res := range attacker.Attack(targeter, rate, duration) {
				window.Add(res)
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		byt, err := json.Marshal(window.Snapshot())
		if err != nil {
			panic(errors.Wrap(err, "Could not marshal window in handlefunc '/'"))
		}

		w.Write(byt)
	})

	http.ListenAndServe(":8081", nil)
}
