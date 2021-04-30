package main

import "fmt"

func main() {
	q := make(chan int)
	cs := send(q)

	receive2(cs, q)

	fmt.Println("exit")
}

func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}

func send(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
