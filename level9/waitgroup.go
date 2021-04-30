package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 并发，执行顺序不可预测
func main() {

	fmt.Println("begin cpus:", runtime.NumCPU())
	fmt.Println("begin gs:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("One")
		wg.Done()
	}()
	go func() {
		fmt.Println("Two")
		wg.Done()
	}()

	go func() {
		fmt.Println("Three")
		wg.Done()
	}()

	fmt.Println("mid cpus:", runtime.NumCPU())
	fmt.Println("mid gs:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("exit")
	fmt.Println("end cpus:", runtime.NumCPU())
	fmt.Println("end gs:", runtime.NumGoroutine())
}
