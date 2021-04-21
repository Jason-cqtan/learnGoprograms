package main

import "fmt"

// 定义struct
type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person{
		first: `jason`,
		last:  `tan`,
		favFlavors: []string{
			`play game`,
			`watching tv`,
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}

	// 以last作为key 将struct 保存到map
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println("\n-------")
	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
