package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 18
	y = "jason"
	z = true
	s := fmt.Sprintf("%v\t %v\t %v",x,y,z)
	fmt.Println(s)
}
