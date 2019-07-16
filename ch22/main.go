package main

import "fmt"

// panic停止当前函数执行（确定是停止当前函数？），一直向上返回，执行每层的recover
// 如果没有遇到recover
// recover只能在defer调用中使用，获取panic值
func main() {
	// 常规使用
	// try0()
	protect(try0)

	// 处理程序意外
	// try1()
	// try2()
	protect(try1)
	protect(try2)
}

func try0() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover说：不让你退出\n\t", e)
		}
	}()

	panic("panic说：我要退出啦")
}

func try1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("try1 to exit")
}

func try2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("try2 to exit")
}

// 装饰器
func protect(g func()) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover\n\t", e)
		}
	}()

	g()
}
