package main

import "fmt"

// 自定义struct
type vehicle struct {
	doors int
	color string
}

// 嵌套自定义struct并扩展字段
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
