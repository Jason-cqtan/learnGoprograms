package main

import "fmt"

func main() {
	c := make(chan int)

	// c<-42
	//fmt.Println(<-c)

	// 使用goroutines
	go func() {
		c <- 27
	}()

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)

	fmt.Println("------")
	// 使用buffer
	c2 := make(chan int, 2)
	c2 <- 26
	c2 <- 29
	fmt.Println(<-c2)

	// comma ok idiom
	v, ok := <-c2
	fmt.Println(v, ok)

	close(c2)

	v, ok = <-c2
	fmt.Println(v, ok)
}
