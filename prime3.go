package main

import (
	"fmt"
	//"math"
)

const total = 1000

func main() {
	//	looptime:=int(math.Sqrt(total))
	//looptime:=2
	var wait [2] chan int
	wait[0]=make(chan int)
	go count2(wait[0])
	for i:=0;i<total;i++{
		prime:=<-wait[0]
		fmt.Println("prime ", i, prime)
		wait[1]=make(chan int)
		go filter2(prime,wait[0],wait[1])
		wait[0]=wait[1]
	}

}

func count2(origin chan int) {

	for n := 2; ; n++ {
		origin <- n
		fmt.Println("chan 1 in", n)
	}
	//	close(origin)

}

func filter2(prime int, seq chan int, wait chan int) {

	for {
		num:=<-seq
		//fmt.Println("filter",prime,num)
		if num%prime != 0 {
			wait <- num
			fmt.Println("IN", num, "with",prime)
		}
	}
}
