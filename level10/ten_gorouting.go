package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// 开10个goroutines，每个goroutine 保存10个数字
	//c := make(chan int)
	//for j := 0; j < 10; j++ {
	//	go func() {
	//		for i := 0; i < 10; i++ {
	//			c <- i
	//		}
	//	}()
	//	fmt.Println("ROUTINES", runtime.NumGoroutine())
	//}
	//
	//for k := 0; k < 100; k++ {
	//	fmt.Println(k, <-c)
	//}
	//fmt.Println("exit")

	//fmt.Println("封装-----")
	x := 10
	y := 10
	d := tenGoroutines(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-d)
	}

	f := tenGroutines2(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-f)
	}


	fmt.Println("ROUTINES", runtime.NumGoroutine())
	fmt.Println("exit")

}

func tenGroutines2(x, y int) <-chan int {
	var wg sync.WaitGroup
	c := make(chan int)
	go func() {
		for i := 0; i < x; i++ {
			wg.Add(1)
			go func(m int) {
				for j := 0; j < y; j++ {
					c <- j
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()
	return c
}

func tenGoroutines(x, y int) <-chan int {
	c := make(chan int)
	for j := 0; j < x; j++ {
		go func() {
			for i := 0; i < y; i++ {
				c <- i
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}
