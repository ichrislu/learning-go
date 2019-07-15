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
}

func computeDiv(d *Divide) (result int, err error) {
	if d.divider == 0 {
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
