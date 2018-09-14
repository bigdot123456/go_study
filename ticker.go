package main

import "time"

func main() {
	// 无法取消
	tick := time.Tick(1 * time.Minute)
	for _ = range tick {
		// do something
	}

	// 可通过调用ticker.Stop取消
	ticker := time.NewTicker(1 * time.Minute)
	for _ = range tick {
		// do something
	}

}
