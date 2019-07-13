package main

import "fmt"

// 切片是数组的连续片断引用，是引用类型
func main() {
	// 初始化[]表示切片类型，不需要指定大小
	// 也可以用make初始化，make([]T, length, capacity)

	ary := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("ary=%v\n", ary)

	a1 := ary[2:6]
	fmt.Printf("a1=%v\n", a1)

	a2 := ary[2:]
	fmt.Printf("a2=%v\n", a2)

	a3 := ary[:6]
	fmt.Printf("a3=%v\n", a3)
	// len为切片的实际长度
	// capacity为切片第1个元素在底层数组的位置，到底层数组结尾位置的长度
	fmt.Printf("len:%d, capacity:%d\n", len(a3), cap(a3))

	a4 := ary[:]
	fmt.Printf("a4=%v\n", a4)

	a5 := a1[1:3]
	fmt.Printf("a5=%v\n", a5)

	// 这里只取到切片的结尾处
	a6 := a3[2:]
	fmt.Printf("a6=%v\n", a6)

	// 这里可超过切片a3的结尾，取到数组的第6个索引位置
	a7 := a3[2:7]
	fmt.Printf("a7=%v\n", a7)

	change(a7)
	fmt.Printf("changed a7=%v\n", a7)
	fmt.Printf("changed a6=%v\n", a6)
	fmt.Printf("changed ary=%v\n", ary)
}

func change(ary []int) {
	ary[0] = 1000
}
