package main

import (
	"fmt"
	"time"
)

func A(c chan int) {

	for i := 0; i < 1024; i++ {

		c <- i
		fmt.Println("insert:", i)
	}

}

func B(c chan int) {

	for val := range c {

		fmt.Println("Value:", val)

	}

}

func main() {

	chs := make(chan int, 1024)

	//只要有通道操作一定要放到goroutine中否则 会堵塞当前的主线程 并且导致程序退出

	//对于同步通道 或者带缓冲的通道 一定要封装成函数 使用 goroutine 包装

	go A(chs)

	go B(chs)

	time.Sleep(1e9)

}
