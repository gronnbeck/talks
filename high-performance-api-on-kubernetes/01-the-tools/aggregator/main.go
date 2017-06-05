package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	serversEnv := os.Getenv("VEGETA_SERVERS")
	runtime := os.Getenv("RUNTIME")

	if serversEnv == "" && runtime != "K8S" {
		panic("You are not running this on k8s. Need to specify servers")
	}

	servers := strings.Split(serversEnv, ",")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.WriteHeader(200)

		if runtime == "K8S" {
			servers = vegetaserversOnKube()
			if len(servers) == 0 {
				log.Println("No servers found. Sleeping 60 seconds before retrying")
				time.Sleep(60 * time.Second)
			}
		}

		nofServers := len(servers)
		done := make(chan vegeta.Metrics)

		for i, server := range servers {
			go func(i int, server string) {
				log.Printf("Fetching data from server: %v", server)
				res, err := http.Get(server)
				if err != nil {
					log.Println(errors.Wrap(err, "Could not get result from server"))
					return
				}
				byt, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Println(errors.Wrap(err, "Could not read payload from body"))
					return
				}
				var metrics vegeta.Metrics
				err = json.Unmarshal(byt, &metrics)
				if err != nil {
					log.Println(errors.Wrap(err, "Could not parse data from payload"))
					return
				}

				done <- metrics
			}(i, server)
		}

		parts := make([]vegeta.Metrics, 0)

		timeout := time.After(1 * time.Second)
		for i := 0; i < nofServers; i++ {
			select {
			case m := <-done:
				parts = append(parts, m)
				continue
			case <-time.After(200 * time.Millisecond):
				log.Println("One of the results timed-out")
			case <-timeout:
				break
			}
		}

		byt, err := json.Marshal(parts)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(byt)
	})

	http.ListenAndServe(":8082", nil)
}
