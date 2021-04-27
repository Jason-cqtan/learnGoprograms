package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		"jason tan",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}

func changeMe(p *person) {
	//p.name = "Miss Moneypenny"
	(*p).name = "Miss Moneypenny"
}
