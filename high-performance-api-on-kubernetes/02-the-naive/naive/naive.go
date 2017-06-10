package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gronnbeck/talks/high-performance-api-on-kubernetes/02-the-naive/domain"
	"github.com/gronnbeck/talks/high-performance-api-on-kubernetes/02-the-naive/naive/setupredis"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/pkg/errors"
)

var (
	redisAddr = flag.String("redis-addr", "localhost:6379", "tcp addr to connect redis tor")
	redisPass = flag.String("redis-pass", "", "password to redis server")
	noValue   = errors.New("No value added yet")
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

type response struct {
	err  error
	resp domain.Response
}

func fetchFromRedis(redis *pool.Pool) (*domain.Response, error) {
	hasCmd := redis.Cmd("HEXISTS", "currency", "current-value")

	if hasCmd.Err != nil {
		return nil, errors.Wrap(hasCmd.Err, "Has command failed")
	}

	res, err := hasCmd.Int()

	if err != nil {
		return nil, errors.Wrap(err, "Could not convert int/bool to int")
	}

	if res == 0 {
		return nil, noValue
	}

	redisResp := redis.Cmd("HGET", "currency", "current-value")

	if redisResp.Err != nil {
		return nil, errors.Wrap(redisResp.Err, "Could not fetch current-value")
	}

	current, err := redisResp.Float64()

	if err != nil {
		return nil, errors.Wrap(err, "Could not convert to current-value Float64")
	}

	return &domain.Response{Current: current}, nil
}

func get(redis *pool.Pool, w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	_ = ctx

	returnChan := make(chan response)
	go func() {
		resp, err := fetchFromRedis(redis)
		if err != nil {
			returnChan <- response{err: err}
		} else {
			returnChan <- response{resp: *resp}
		}
	}()

	select {
	case res := <-returnChan:
		if res.err != nil {
			errorStr := res.err.Error()
			w.Write(domain.Response{Error: &errorStr}.JSON())
			return
		}
		w.Write(res.resp.JSON())
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			errStr := context.DeadlineExceeded.Error()
			w.WriteHeader(408)
			w.Write(domain.Response{Error: &errStr}.JSON())
		}
	}

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
