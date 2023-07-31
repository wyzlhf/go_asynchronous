package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res

}

func Barrier(urls ...string) {
	requestNumber := len(urls)
	in := make(chan BarrierResponse, requestNumber)
	response := make([]BarrierResponse, requestNumber)
	defer close(in)
	for _, url := range urls {
		go doRequest(in, url)
	}
	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR :", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)
		}
	}
}
func main() {
	Barrier([]string{
		"https://www.baidu.com",
		"https://www.weibo.com",
		"https://www.shirdon.com",
	}...)
}
