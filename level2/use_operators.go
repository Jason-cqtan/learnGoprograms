package main

import "fmt"

// 使用操作符== <= >= != < >
func main() {
	a := (18 == 18)
	b := (12 <= 18)
	c := (18 >= 27)
	d := (23 != 27)
	e := (18 < 27)
	f := (18 > 27)

	fmt.Println(a, b, c, d, e, f)
}
