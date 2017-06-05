package main

import "encoding/json"

type response struct {
	Error   *string `json:"error,omitempty"`
	Current float64 `json:"current,omitempty"`
}

func (r response) json() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return byt
}
