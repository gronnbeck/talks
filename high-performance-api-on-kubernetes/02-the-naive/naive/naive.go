package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"./setupredis"
	"github.com/mediocregopher/radix.v2/pool"
)

var (
	redisAddr = flag.String("redis-addr", "localhost:6379", "tcp addr to connect redis tor")
	redisPass = flag.String("redis-pass", "", "password to redis server")
	noValue   = "No value added yet"
)

func init() {
	flag.Parse()

	if redisAddr == nil || *redisAddr == "" {
		panic("Reids addr cannot be empty")
	}
}

func main() {
	redis, err := setupredis.NewWait(*redisAddr, *redisPass)

	if err != nil {
		panic("Couldn't setup redis")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.ToLower(r.Method) == "get" {
			get(redis, w, r)
		}
		if strings.ToLower(r.Method) == "post" {
			post(redis, w, r)
		}

	})

	log.Println("HTTP server running: 8080")
	http.ListenAndServe(":8080", nil)
}

func get(redis *pool.Pool, w http.ResponseWriter, r *http.Request) {
	hasCmd := redis.Cmd("HEXISTS", "currency", "current-value")

	if hasCmd.Err != nil {
		panic(hasCmd.Err)
	}

	res, err := hasCmd.Int()

	if err != nil {
		panic(err)
	}

	if res == 0 {
		w.Write(response{Error: &noValue}.json())
		return
	}

	redisResp := redis.Cmd("HGET", "currency", "current-value")

	if redisResp.Err != nil {
		panic(err)
	}

	current, err := redisResp.Float64()

	if err != nil {
		panic(err)
	}

	w.Write(response{Current: current}.json())
}

func post(redis *pool.Pool, w http.ResponseWriter, r *http.Request) {
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req request
	unmarshallErr := json.Unmarshal(byt, &req)
	if unmarshallErr != nil {
		panic(err)
	}
	res := redis.Cmd("HSET", "currency", "current-value", req.Value)
	if res.Err != nil {
		panic(res.Err)
	}

}
