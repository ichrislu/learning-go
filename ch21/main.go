package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("测试自定义错误")
	fmt.Println(err)

	err2 := fmt.Errorf("自定义错误，来自:%s", err.Error())
	fmt.Println(err2)

	fmt.Println("===================")

	de := Divide{100, 0}
	if result, err := computeDiv(&de); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("===================")

	// defer在方法调用结束后，return调用前执行
	// 多个defer同时存在时，是以栈的形式，先进后出，FILO
	// 多用于释放资源（个人理解应该是确保释放资源，因为资源不用应该及时释放，而不是依赖于defer，defer的存在只是为了确保方法完成后资源会释放）
}

func computeDiv(d *Divide) (result int, err error) {
	if d.divider == 0 {
		// 因为Divide实现了Error方法，也就实现了error
		err = d
	} else {
		result = d.dividee / d.divider
	}

	return
}

// 定义除法结构休
type Divide struct {
	dividee int
	divider int
}

// 为除法添加一个Error
func (d Divide) Error() string {
	sFormat := `错误的除法运算
	dividee: %d
	divider: %d
	`

	return fmt.Sprintf(sFormat, d.dividee, d.divider)
}
