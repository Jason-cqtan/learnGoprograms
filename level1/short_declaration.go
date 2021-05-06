package main

import "fmt"

func main() {
	x := 18
	y := `jason`
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	whoAmI(19.00)
}

func whoAmI(a interface{}) {
	switch t := a.(type) {
	case bool:
		fmt.Println(`bool`)
	case string:
		fmt.Println(`string`)
	case int:
		fmt.Println(`int`)
	default:
		fmt.Printf("dont know type %T \n", t)
	}
}
