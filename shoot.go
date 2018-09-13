package main

import (
	"fmt"
	"time"
)


func fixedShooting(msgChan chan string) {
	var times = 3
	var t = 1
	for {
		if t <= times {
			msgChan <- "fixed shooting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func threePointShooting(msgChan chan string) {
	var times = 5
	var t = 1
	for {
		if t <= times {
			msgChan <- "three point shooting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func main() {
	cFixeds := make(chan string)
	c3Points := make(chan string)

	go fixedShooting(cFixeds)
	go threePointShooting(c3Points)

	go func() {
		for {
			select {
			case msg1 := <-cFixeds:
				fmt.Println(msg1)
			case msg2 := <-c3Points:
				fmt.Println(msg2)
			case <-time.After(time.Second * 5):
				fmt.Println("timeout, check again...")
			}
		}

	}()

	var input string
	fmt.Scanln(&input)
}