package main

import (
	"fmt"
)

func f1(in chan int) {
	a := <-in
	fmt.Println("run", a)

}

func main() {
	out := make(chan int)
	go f1(out)
	out <- 2
	fmt.Println("hello!")

}
