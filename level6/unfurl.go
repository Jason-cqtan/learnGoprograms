package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := sum(ii...)
	fmt.Println(n)

	n2 := sum2(ii)
	fmt.Println(n2)
}

// ... 展开切片
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sum2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
