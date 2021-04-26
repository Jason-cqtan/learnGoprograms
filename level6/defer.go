package main

import "fmt"

func main() {
	defer foo1()
	fmt.Println("main")

	//main; foo1 in; foo1 defer
}

func foo1() {
	defer func() {
		fmt.Println("foo1 defer")
	}()

	fmt.Println("foo1 in")
}
