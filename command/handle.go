package command

import (
	ht "Serpent/http"
	"flag"
	"fmt"
)

//func Usage() {
//	//   _____                            _
//	//  / ____|                          | |
//	// | (___   ___ _ __ _ __   ___ _ __ | |_
//	//  \___ \ / _ | '__| '_ \ / _ | '_ \| __|
//	//  ____) |  __| |  | |_) |  __| | | | |_
//	// |_____/ \___|_|  | .__/ \___|_| |_|\__|
//	//                  | |
//	//                  |_|
//	fmt.Println("   _____                            _   \n  / ____|                          | |  \n | (___   ___ _ __ _ __   ___ _ __ | |_ \n  \\___ \\ / _ | '__| '_ \\ / _ | '_ \\| __|\n  ____) |  __| |  | |_) |  __| | | | |_ \n |_____/ \\___|_|  | .__/ \\___|_| |_|\\__|\n                  | |                   \n                  |_|                   ")
//	fmt.Println("Options:\n")
//	flag.PrintDefaults()
//
//}
func httpflag(args []string) {
	//http功能
	http := flag.NewFlagSet("http", flag.ExitOnError)
	//定义help内容
	http.Usage = func() {
		fmt.Println("   _____                            _   \n  / ____|                          | |  \n | (___   ___ _ __ _ __   ___ _ __ | |_ \n  \\___ \\ / _ | '__| '_ \\ / _ | '_ \\| __|\n  ____) |  __| |  | |_) |  __| | | | |_ \n |_____/ \\___|_|  | .__/ \\___|_| |_|\\__|\n                  | |                   \n                  |_|                   ")
		fmt.Println("Options:\n")
		http.PrintDefaults()
	}
	//创建参数
	//url
	var u = http.String("u", "", "url: url address")

	//测压线程
	var c = http.Int("c", 0, "concurrent: concurrent number")

	//请求类型
	var t = http.String("t", "get", "type: request type, 'get' or 'post'")
	//获取参数
	http.Parse(args)
	//执行
	//初始化返回值
	var RequestQuantity int64 = 0
	var Success int64 = 0
	//传值
	url := *u
	concurrent := *c
	requesttype := *t
	ht.Request(url, concurrent, requesttype, &RequestQuantity, &Success)
	for true {
		fmt.Printf("\rRequestQuantity: %d,Success: %d", RequestQuantity, Success)
	}
}
