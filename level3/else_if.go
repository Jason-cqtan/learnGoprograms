package main

import "fmt"

func main() {
	x := "jason"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "jason" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
