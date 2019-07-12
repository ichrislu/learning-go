package main

import "fmt"

func main() {
	a := 1
	a = a << 3 // 乘以2的N次方
	fmt.Printf("a=%d\n", a)
	a = a >> 3 // 除以2的N次方
	fmt.Printf("a=%d\n", a)
}
