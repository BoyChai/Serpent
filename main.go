package main

import (
	"Serpent/control"
	http2 "Serpent/http"
	"flag"
	"fmt"
)

//url
var u = flag.String("u", "", "url: url address")

//测压线程
var c = flag.Int("c", 0, "concurrent: concurrent number")

//请求类型
var t = flag.String("t", "", "type: request type, 'get' or 'post'")

//主线程
func main() {

	//解析命令行参数
	flag.Parse()
	//初始化
	var RequestQuantity int64 = 0
	var Success int64 = 0
	//传值
	url := *u
	concurrent := *c
	requesttype := *t
	http2.Request(url, concurrent, requesttype, &RequestQuantity, &Success)
	for true {
		fmt.Printf("\rRequestQuantity: %d,Success: %d", RequestQuantity, Success)
	}
	control.Wg.Wait()
}
