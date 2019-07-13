package main

import "fmt"

func main() {
	// [3]，中括号中有数字代表数组类型

	// 指定长度数组，索引位置未赋初值则为数据类型默认值
	ary1 := [3]int{1, 2}
	fmt.Println(ary1)
	fmt.Println(ary1[1])

	// 数组赋值时确定长度，...做为数组不可少，缺少时代表切片
	ary2 := [...]string{"chris", "emily", "shmily", "coco"}
	fmt.Println(ary2)
	ary2[2] = "coby"
	fmt.Println(ary2)

	for i := 0; i < len(ary2); i++ {
		fmt.Printf("%s,", ary2[i])
	}

	fmt.Println()

	fmt.Println("取索引和值")
	// 取索引和值
	for idx, val := range ary2 {
		fmt.Printf("%d:%s\n", idx, val)
	}

	fmt.Println("取索引")
	// 取索引
	for idx := range ary2 {
		fmt.Printf("%d\n", idx)
	}

	fmt.Println("取值")
	// 取值
	for _, val := range ary2 {
		fmt.Printf("%s\n", val)
	}

	// 普通参数
	ary3 := [3]int{1, 2, 3}
	change(ary3)
	fmt.Print("outside:")
	fmt.Println(ary3)

	// 指针参数
	fmt.Println()
	change2(&ary3)
	fmt.Print("outside:")
	fmt.Println(ary3)
}

// 复制了一份数据，所以受影响范围只在函数内
func change(ary [3]int) {
	ary[2] = 100
	fmt.Print("in func:")
	fmt.Println(ary)
}

// 参数接收一个指针类型，与调用处指向同一块内在，所以修改这里，影响外面
func change2(ary *[3]int) {
	// 此处的ary[2]不需要加*，go可以自动识别
	ary[2] = 1000
	fmt.Print("in fun2:")
	fmt.Println(ary)
}
