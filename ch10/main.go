package main

import (
	"fmt"
	"math"
)

func main() {
	// 异常处理
	if val, err := compute(1, 2, "++"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("val:%d\n", val)
	}

	// 多返回值
	a, b := 1, 2
	x, y := swap(a, b)
	fmt.Printf("x=%d, y=%d\n", x, y)

	// 函数参数
	v := com(math.Max, 0.11, 0.2)
	fmt.Printf("v=%f\n", v)
}

// op为变量：即op
// 从func弄好到逗号前的参数为类型，如 func(float64, float64) float64
func com(op func(float64, float64) float64, a, b float64) float64 {
	return op(a, b)
}

func compute(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsuppert operate:%s", op)
		// panic(fmt.Sprintf("unsuppert operate:%s", op))
	}
}

func swap(a, b int) (x, y int) {
	return b, a
}
