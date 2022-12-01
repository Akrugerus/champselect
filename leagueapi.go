package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var base_url = os.Getenv("RIOT_URL")

func api(method string, url string, body io.Reader) (data []byte, err error) {
	res, err := http.NewRequest(method, base_url+url, body)
	if err != nil {
		return nil, err
	}
	m, err := json.Marshal(res.Body)
	if err != nil {
		return nil, err
	}
	return m, nil
}
