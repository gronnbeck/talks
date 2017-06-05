package domain

import "encoding/json"

type Response struct {
	Error   *string `json:"error,omitempty"`
	Current float64 `json:"current,omitempty"`
}

func (r Response) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return byt
}
