package main

import "fmt"

func main() {
	// 匿名函数，n 的阶乘
	factorial := func(n int) int {
		total := 1
		for ; n > 0; n-- {
			fmt.Println(n)
			total *= n
		}
		return total
	}

	fmt.Printf("%T\n", factorial)
	fmt.Println(factorial(5))

	fmt.Println("done")
}
