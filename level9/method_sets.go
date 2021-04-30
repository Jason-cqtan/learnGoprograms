package main

import "fmt"

type person struct {
	First string
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p *person) speak() {
	fmt.Println("hi, Im ", p.First)
}

func main() {
	p := person{
		"jason",
	}

	saySomething(&p)
	p.speak()

	//saySomething(p)
}
