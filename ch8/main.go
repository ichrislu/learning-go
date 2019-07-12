package main

import "fmt"

func main() {
	// switch语句，case后面不用加break
	// 默认case匹配成功后，不地执行其他case，如果需要执行后面case，可以使用fallthrough
	a := 2

	// switch方式1
	switch {
	case a == 1:
		fmt.Println("a=1")
	case a == 2:
		fmt.Println("a=2")
		fallthrough // 此时，会进行下一项case，且不做条件判断！！！
	case a == 3:
		fmt.Println("a=3")
	case a == 4:
		fmt.Println("a=4")
	case a == 5:
		fmt.Println("a=5")
	default:
		fmt.Println("default")
	}

	fmt.Println("==============")

	// switch方式2
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2, 3, 4:
		fmt.Println("a in 2, 3 ,4")
		fallthrough
	case 5:
		fmt.Println("a=5")
	default:
		fmt.Println("default")
	}
}
