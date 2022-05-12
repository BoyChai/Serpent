package command

import (
	ht "Serpent/ability/http"
	sniff2 "Serpent/ability/sniff"
	"Serpent/command/judge"
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

//func blastflag(args []string) {
//	//爆破功能添加
//	blast := flag.NewFlagSet("blast", flag.ExitOnError)
//	//定义help内容
//	blast.Usage = func() {
//		fmt.Println("   _____                            _   \n  / ____|                          | |  \n | (___   ___ _ __ _ __   ___ _ __ | |_ \n  \\___ \\ / _ | '__| '_ \\ / _ | '_ \\| __|\n  ____) |  __| |  | |_) |  __| | | | |_ \n |_____/ \\___|_|  | .__/ \\___|_| |_|\\__|\n                  | |                   \n                  |_|                   ")
//		fmt.Println("Options:\n")
//		blast.PrintDefaults()
//	}
//	//指定类型
//	//var blasttype = blast.String("t", "", "blasttype: Blasting mode")
//	////指定目标
//	//var host = blast.String("h", "", "host: Specify host")
//	////指定端口
//	//var prot = blast.Int("p", 0, "prot: Specify host port")
//	////指定用户
//	//var user = blast.String("u", "", "user: Specify blasting user")
//	////指定密码文件
//	//var passwordfile = blast.String("f", "", "passwordfile: Specify password file")
//	//获取参数
//	blast.Parse(args)
//}
func sniffflag(args []string) {

	//嗅探功能
	sniff := flag.NewFlagSet("sniff", flag.ExitOnError)
	//定义help内容
	sniff.Usage = func() {
		fmt.Println("   _____                            _   \n  / ____|                          | |  \n | (___   ___ _ __ _ __   ___ _ __ | |_ \n  \\___ \\ / _ | '__| '_ \\ / _ | '_ \\| __|\n  ____) |  __| |  | |_) |  __| | | | |_ \n |_____/ \\___|_|  | .__/ \\___|_| |_|\\__|\n                  | |                   \n                  |_|                   ")
		fmt.Println("Options:\n")
		sniff.PrintDefaults()
	}
	//var posttype = sniff.String("t", "tcp", "posttype: Specify sniff type")
	var host = sniff.String("ip", "", " host IP: Specify sniff host")
	var port = sniff.String("p", "", "port: Specify sniff port")
	var thread = sniff.Int("c", 2048, "Cache thread channel: Specifies the number of cache channels,Not too high")
	sniff.Parse(args)
	ports, err := judge.PortConversion(*port)
	if err != nil {
		fmt.Println("[Serpent]: ", err)
		return
	}
	sniff2.RunTCPSniff(*host, ports, *thread)

}
