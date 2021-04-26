package main

import "fmt"

func main() {
	num := proxy()()
	fmt.Println(num)
}

// 函数返回函数类型
func proxy() func() int {
	return func() int {
		return 27
	}
}
