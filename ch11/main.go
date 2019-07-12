package main

import "fmt"

// go语言只有值传递，没有引用传递
// 指针变量指向一个值的内存地址，*用于指定变量是作为一个指针
// 即：类型前面带*，表明该类型的变量为指针？

// &符号用于取后面变量的内存地址
// *符号用于取变量值
func main() {
	var p = 30  // 普通变量
	var ip *int // 指针变量
	ip = &p     // 取普通变量的内存地址，赋值给指针变量
	*ip = 20    // 指针变量赋值，因为指向了普通变量，所以p值也改了
	fmt.Printf("p=%d\n", p)

	valRef(p)
	fmt.Printf("p=%d\n", p)

	valRefPtr(&p) // 取p的地址当指针参数
	fmt.Printf("p=%d\n", p)
}

// 参数copy一份p的值
func valRef(p int) {
	p = 1000
}

// 要求参数是指针类型
func valRefPtr(p *int) {
	// 设置该指针就是的值为1001
	*p = 1001
}
