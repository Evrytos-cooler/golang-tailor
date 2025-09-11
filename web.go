package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// GO 多线程 + 阻滞调用 JS 单线程事件循环 + 非阻塞 I/O
// 对比 	成本		切换成本     	调度方式    	栈管理 		通信方式
// 协程 	~2KB    	用户态（低） 	GO运行时调度 	动态增长 	channel > 锁 + 共享内存
// 线程 	1-8MB		内核态（高）	操作系统		固定大小	锁 / 信号量 > 消息管道
func WebFetch() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
	// 捕获错误并将错误输出到标准错误流并调用终止程序 (os.Exit(1))
	// 只有当 http.listenAndServe 返回的时候才会执行, 换句话说 http.listenAndServe 是注册的
	// 如果我们希望有非阻塞当前代码块的效果，需要使用 goroutine 让他去别的协程执行， 而不是在当前协程
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// 这里的内容正常情况下不会被调用，这和 JS 的单线程非阻滞 IO 有本质区别
}

func handler(write http.ResponseWriter, request *http.Request) {
	url := request.Host
	path := request.URL.Path
	fmt.Fprintf(write, "URL = %s%s\n", url, path)
	mu.Lock()
	count++
	mu.Unlock()
}

func countHandler(write http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(write, "path / has been request %d times \n", count)
	mu.Unlock()
}
