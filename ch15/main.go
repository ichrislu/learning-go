package main

import "fmt"

func main() {
	// 仅声明一个map类型变量
	var m1 map[string]int
	fmt.Println(m1, m1 == nil)

	// 初始化map变量，才能做put操作
	m1 = make(map[string]int)
	m1["chris"] = 97

	fmt.Println(m1)

	fmt.Println()

	// 定义+直接初始化
	m2 := map[string]int{
		"chris": 1,
		"una":   2,
		"test":  3,
	}
	fmt.Println(m2)
	fmt.Println(m2["chris"])

	for key, value := range m2 {
		fmt.Printf("%s=%d\n", key, value)
	}

	delete(m2, "test")
	fmt.Println(m2)
}
