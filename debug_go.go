package main

import (
	"fmt"
	"runtime"
	"time"
)

func f3() {
	for {
		fmt.Println("11111")
		time.Sleep(2*time.Second)
	}
}

func f2() {
	for {
		fmt.Println("22222")
		time.Sleep(3*time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go f3()
	go f2()

	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}

}
