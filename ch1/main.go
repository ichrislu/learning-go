package main

import "fmt"

var (
	a1 = 1
	a2 = "golang"
)

func main() {
	fmt.Printf("hello world!\n")

	fmt.Printf("a1:%d, a2:%s\n", a1, a2)

	// bool类型用%t
	b1, b2 := "b1", true
	fmt.Printf("b1:%s, b2:%t", b1, b2)
}
