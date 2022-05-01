package http

import (
	"Serpent/control"
	"errors"
	"net/http"
	"strings"
)

//发起请求
func Request(url string, Concurrent int, RequestType string, RequestQuantity *int64, Success *int64) error {
	switch RequestType {
	case "get":
		for i := 1; i <= Concurrent; i++ {
			control.Wg.Add(1)
			go getUrl(url, &RequestQuantity, &Success)
		}
		return nil
	case "post":
		for i := 1; i <= Concurrent; i++ {
			control.Wg.Add(1)
			go postUrl(url, &RequestQuantity, &Success)
		}
		return nil
	default:
		return errors.New("Request type should be 'get' or 'post'")
	}
}

//get请求
func getUrl(url string, RequestQuantity **int64, Success **int64) {
	for true {
		resp, _ := http.Get(url)
		**RequestQuantity++

		if resp != nil {
			if resp.StatusCode < 300 && resp.StatusCode > 199 {
				**Success++
			}
			resp.Body.Close()
		}
	}

	control.Wg.Done()
}

//post请求
func postUrl(url string, RequestQuantity **int64, Success **int64) {
	for true {
		resp, _ := http.Post(url, "application/json;charset=utf-8", strings.NewReader(string([]byte("{'':''}"))))
		**RequestQuantity++
		if resp != nil {
			resp.Body.Close()
			if resp.StatusCode < 300 && resp.StatusCode > 199 {
				**Success++
			}
		}
	}
	control.Wg.Done()
}
