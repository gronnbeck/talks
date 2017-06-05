package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gronnbeck/talks/high-performance-api-on-kubernetes/01-the-tools/vegetaserver/rollingwindow"

	"github.com/pkg/errors"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	targetURL := os.Getenv("TARGET_URL")

	if targetURL == "" {
		panic("Must specify TARGET_URL")
	}

	rate := uint64(1000)
	rateEnv := os.Getenv("ATTACK_RATE")
	if rateEnv != "" {
		rateInt, err := strconv.Atoi(rateEnv)
		if err != nil {
			log.Printf("Could not convert %v to an integer. Using default value 1000", rateInt)
		} else {
			rate = uint64(rateInt)
		}

	}

	window := rollingwindow.New(60)

	go func() {
		for {
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
		byt, err := json.Marshal(window.Snapshot(time.After(200 * time.Millisecond)))
		if err != nil {
			panic(errors.Wrap(err, "Could not marshal window in handlefunc '/'"))
		}

		w.Write(byt)
	})

	http.ListenAndServe(":8081", nil)
}
