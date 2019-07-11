package main

import "fmt"

// 常量定义不使用，不报错
const str1 = "str1 value"

var str2 = "str2 value"

func main() {
	fmt.Printf("xx\n")
	fmt.Println(str2)

	type myInt = int32
	var i1 myInt = 1
	fmt.Printf("i1:%d, %T", i1, i1)
}
