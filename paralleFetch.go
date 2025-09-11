package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// main 函数本身也是一个 goroutine
func ParallelFetch(url []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range url {
		// go function() 是非阻塞的，发起 goroutine 之后就马上继续当前线程
		go fetch(url, ch, start)
	}

	for range url {
		// 但是管道通信是会阻塞的，发起之后会等待下次到达
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f elapsed \n ", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, start time.Time) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//ioutil 可以看做一个垃圾桶， io.Copy 将 B copy 到 A 上
	bytesLength, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("从 %s 读取时出错  %v", url, err)
		return
	}
	cost := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s ", cost, bytesLength, url)
}
