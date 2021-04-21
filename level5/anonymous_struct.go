package main

import "fmt"

func main() {

	// 匿名struct
	p := struct {
		first     string
		favDrinks []string
		age       int
	}{
		first: "jason",
		favDrinks: []string{
			`Martini`,
			`Water`,
		},
		age: 27,
	}

	fmt.Println(p.first)
	fmt.Println(p.favDrinks)
	fmt.Println(p.age)

	for i, v := range p.favDrinks {
		fmt.Println(i, v)
	}
}
