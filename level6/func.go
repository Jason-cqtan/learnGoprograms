package main

import "fmt"

func main() {
	n := foo()

	num, s := bar()
	fmt.Println(n)
	fmt.Println(num, s)
}

// func receiver, identifier, params, returns
func foo() int {
	return 27
}

// 函数返回多个值
func bar() (int, string) {
	return 1993, "big brother is Watching"
}
