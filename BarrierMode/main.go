package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

type BarrierResponse struct {
	Err    error
	Resp   string
	Status int
}

//构造请求

func doRequest(out chan<- BarrierResponse, url string) {
	res := BarrierResponse{}
	client := http.Client{
		Timeout: time.Duration(20 * time.Second),
	}
	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
}
