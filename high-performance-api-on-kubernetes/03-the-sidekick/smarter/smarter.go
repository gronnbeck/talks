package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gronnbeck/talks/high-performance-api-on-kubernetes/02-the-naive/domain"
	"github.com/gronnbeck/talks/high-performance-api-on-kubernetes/02-the-naive/naive/setupredis"
	"github.com/mediocregopher/radix.v2/pool"
)

var (
	readRedisAddr = flag.String("read-redis-addr", "localhost:6379", "tcp addr to connect redis tor")
	readRedisPass = flag.String("read-redis-pass", "", "password to redis server")

	writeRedisAddr = flag.String("write-redis-addr", "localhost:6379", "tcp addr to connect redis tor")
	writeRedisPass = flag.String("write-redis-pass", "", "password to redis server")

	noValue = "No value added yet"
)

func init() {
	flag.Parse()

	if readRedisAddr == nil || *readRedisAddr == "" {
		panic("Read Reids addr cannot be empty")
	}

	if writeRedisAddr == nil || *writeRedisAddr == "" {
		panic("Write Reids addr cannot be empty")
	}
}

func main() {
	readRedis, err := setupredis.NewWait(*readRedisAddr, *readRedisPass)

	if err != nil {
		panic("Couldn't setup read redis")
	}

	writeRedis, err := setupredis.NewWait(*writeRedisAddr, *writeRedisPass)

	if err != nil {
		panic("Couldn't setup write redis")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.ToLower(r.Method) == "get" {
			get(readRedis, w, r)
		}
		if strings.ToLower(r.Method) == "post" {
			post(writeRedis, w, r)
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
		w.Write(domain.Response{Error: &noValue}.JSON())
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

	w.Write(domain.Response{Current: current}.JSON())
}

func post(redis *pool.Pool, w http.ResponseWriter, r *http.Request) {
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req domain.Request
	unmarshallErr := json.Unmarshal(byt, &req)
	if unmarshallErr != nil {
		panic(err)
	}
	res := redis.Cmd("HSET", "currency", "current-value", req.Value)
	if res.Err != nil {
		panic(res.Err)
	}

}
