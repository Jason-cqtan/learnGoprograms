package main

// 使用iota
import "fmt"

const (
	year = 2021 + iota
	nextYear
	c
	d
)

func main() {
	fmt.Println(year)
	fmt.Println(nextYear)
	fmt.Println(c)
	fmt.Println(d)
}
