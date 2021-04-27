package main

import "fmt"

// https://golang.org/ref/spec#Types
type cool int

var jason cool
var age int

func main() {
	fmt.Println(jason)
	fmt.Printf("%T\n", jason)

	jason = 18
	fmt.Println(jason)

	age = int(jason)
	fmt.Println(age)
	fmt.Printf("%T\n", age)
}
