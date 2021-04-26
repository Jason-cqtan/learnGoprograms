package main

import "fmt"

func main() {
	f := increasing()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("------")


	g := increasing()
	fmt.Println(g())
	fmt.Println(g())
}

func increasing() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
