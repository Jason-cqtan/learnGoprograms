package main

// 创建无符号和符号常量
import (
	"fmt"
	"testing"
	"time"
)

const (
	a     = 43
	b int = 27
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkHello2(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

func main() {
	fmt.Println(a, b)
	fmt.Println(time.Now())
	fmt.Println("Hello"[1])
	//threeFive()
	//stop := time.After(3 * time.Second)
	//tick := time.NewTicker(1 * time.Second)
	//defer tick.Stop()
	//for {
	//	select {
	//	case <-tick.C:
	//		fmt.Println(time.Now())
	//	case <-stop:
	//		return
	//	}
	//}


}

/**
1000 以下3或5的倍数的和
 */
func threeFive() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}