package main

import "fmt"

func main() {
	cs := make(chan<- int)

	cr := make(<-chan int)

	go func() {
		//send to receive-only type <-chan int
		//cr<- 12
		cs <- 32
	}()

	//receive from send-only type chan<-
	//fmt.Println(<-cs)

	//fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

}
