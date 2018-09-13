package main

import (
	"fmt"
)

const total = 1000

func main() {
	origin := make(chan int)
	wait := make(chan int)
	go count(origin)
	go filter(2, origin, wait)
	for i:= range wait {
		fmt.Println("chan ", i, <-wait)
	}
	//<-wait
	//defer close(origin)
}

func count(origin chan int) {

	for n := 2;n<total ; n++ {
		origin <- n
		fmt.Println("chan origin in ",n)
	}
	close(origin)

}

func filter(prime int, seq chan int, wait chan int) {

	for num := range seq {
		if num%prime != 0 {
			wait <- num
			fmt.Println("data", num)
		}
	}
	close(wait)

}
