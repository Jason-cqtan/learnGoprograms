package main

import "fmt"

func main() {
	q := make(chan bool)
	odd := make(chan int)
	even := make(chan int)

	// 发送
	send(even, odd, q)

	// 接收
	receive2(even, odd, q)

	fmt.Println("exit")
}

func receive2(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println(`偶数`, v)
		case v := <-odd:
			fmt.Println(`奇数`, v)
		case <-quit:
			return
		}
	}
}

func send(even, odd chan<- int, q chan<- bool) {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		q <- true
		close(even)
		close(odd)
	}()
}
