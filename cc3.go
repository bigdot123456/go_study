package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)

	for {

		///不停向channel中写入 0 或者1

		select {

		case ch <- 0:

		case ch <- 1:

		}

		//从通道中取出数据

		i := <-ch

		fmt.Println("Value received:", i)

		time.Sleep(1e8)

	}

}
