package main

import (
	"fmt"
	//"math"
)

const total = 1000

func main() {
//	looptime:=int(math.Sqrt(total))
	//looptime:=2
	origin := make(chan int)
	var wait [40] chan int
	wait[0]=make(chan int)
	wait[1]=make(chan int)
	wait[2]=make(chan int)
	go count1(origin)
	go filter1(2, origin, wait[0])
	go filter1(3, wait[0], wait[1])
	go filter1(5, wait[1], wait[2])

	for i := 0; i < 20; i++ {
	//	fmt.Println("prime ", i, <-wait[0])
		fmt.Println("prime ", i, <-wait[2])
	}

	//for i:=2;i<looptime;i++ {
	//	prime:=3//<-wait[i-1]
	//	fmt.Println("prime", i,prime)
	//	go filter1(prime, wait[i-1], wait[i])
	//
	//}
	//
	//for i := 0; i < 40; i++ {
	//	fmt.Println("prime ", i, <-wait[looptime-1])
	//}
	////<-wait
	defer close(origin)
	defer close(wait[0])
	defer close(wait[1])
}

func count1(origin chan int) {

	for n := 2; ; n++ {
		origin <- n
		fmt.Println("chan 1 in", n)
	}
	//	close(origin)

}

func filter1(prime int, seq chan int, wait chan int) {

	for {
		num:=<-seq
		fmt.Println("filter",prime,num)
		if num%prime != 0 {
			wait <- num
			fmt.Println("chanA", prime, "in", num)
		}
	}
}
