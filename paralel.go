package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const START = 1
const LEN = 100
const STEP = 1000

var sum int64
var mu sync.Mutex

func main() {
	fmt.Printf("%d\n", (int64(LEN*STEP) * (int64(LEN*STEP) + 1) / int64(2)))

	sum = 0
	start := START
	for i := 1; i <= STEP; i++ {
		go subsum(start, LEN)
		start += LEN
	}

	for runtime.NumGoroutine() > 1 {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("%d\n", sum)
}

func subsum(start int, len int) {
	var ssum int64
	ssum = 0
	for i := 1; i <= len; i++ {
		ssum += int64(start)
		start++
	}
	mu.Lock()
	sum += ssum
	mu.Unlock()
}
