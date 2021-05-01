package main

import "fmt"

type customErr struct {
	info string
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is this error:%v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
