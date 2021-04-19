package main

import "fmt"

// 打印一个数十进制、二进制、十六进制
func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}
