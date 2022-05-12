package sniff

import (
	"fmt"
	"net"
	"sort"
)

//工作缓存线程
func work(host string, port chan int, results chan int) {
	//读取端口
	for p := range port {
		//组装
		address := fmt.Sprintf("%s:%d", host, p)
		//访问
		conn, err := net.Dial("tcp", address)
		//if p == 7000 || p == 7001 || p == 7002 {
		if p == 7000 {
			fmt.Println(p)
			fmt.Println(err)
		}
		//不成功跳出循环等待下一个端口并回传0，成功则关闭链接返回成功的端口
		if err != nil {
			results <- 0
			continue
		}

		//defer conn.Close()
		conn.Close()
		results <- p
	}

}
func RunTCPSniff(host string, ports []int, thread int) {
	////定时器
	//start := time.Now()

	//创建端口通道并初始化，存有缓存线程数
	port := make(chan int, thread)
	//定义返回端口的通道
	results := make(chan int)
	//返回端口存储在这个切片里
	var openports []int
	//开始创建缓存线程
	for i := 0; i < cap(port); i++ {
		go work(host, port, results)
	}
	//向缓存线程里输入值
	go func() {
		for _, i := range ports {
			port <- i
		}
	}()
	i := 0
	for range ports {
		i++
		post := <-results
		if post != 0 {
			openports = append(openports, post)
		}
	}
	close(port)
	close(results)
	sort.Ints(openports)
	fmt.Println(i)
	for _, p := range openports {
		fmt.Println("打开的端口", p)

	}
	//cost := time.Since(start)
	//fmt.Printf("Run time is [%s]", cost)
}
