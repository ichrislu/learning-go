package main

import "fmt"

// 常量定义不使用，不报错
const str1 = "str1 value"

var str2 = "str2 value"

func main() {
	fmt.Printf("xx\n")
	fmt.Println(str2)

	// 无等号，是类型定义；有等号，是类型别名
	type myInt1 int32   // 定义一个基于int32的新类型，一般用于提高代码可读性
	type myInt2 = int32 // 与int32完全一样，类型别名，一般用于代码兼容性，如换了包名
	var i1 myInt2 = 1
	fmt.Printf("i1:%d, %T", i1, i1)
}
