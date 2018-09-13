package main

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}

}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}

}

func main() {
	ch := make(chan int)
	go Generate(ch)

	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

}

// 从这段代码可以体会到，chan 作为参数的时候传递的是 chan 的位置，chan 被赋值的时候也是改变的 chan 指向的位置，这个有点像是指针

// 也就是说 chan 被赋值不会关闭已经打开的 chan，只是让 chan 指向了一个新的位置。
