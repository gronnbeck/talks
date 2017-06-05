package domain

import "encoding/json"

type Request struct {
	Value float64 `json:"value"`
}

func (r Request) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return byt
}
