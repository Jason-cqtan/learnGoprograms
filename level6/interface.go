package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

// 圆面积 pi *r*r
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 长方形面积 长乘宽
func (s square) area() float64 {
	return s.length * s.width
}

// 拥有面积的接口
type shape interface {
	area() float64
}

// 型状详细
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
		width:  12,
	}

	info(circ)
	info(squa)
}
