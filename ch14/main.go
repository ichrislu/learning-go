package main

import "fmt"

func main() {
	ary := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := ary[2:6]
	fmt.Printf("s1=%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	// 切片追加一个元素，影响切片长度，以及源数组内容（因为是指针）
	// s1 = append(s1, 10)
	// fmt.Printf("s1=%v, len=%d, cap=%d, ary=%v\n", s1, len(s1), cap(s1), ary)

	// 切片追加元素超过切片容量，切片会自动扩容
	s1 = append(s1, 10, 20, 30, 40, 50)
	fmt.Printf("s1=%v, len=%d, cap=%d, ary=%v\n", s1, len(s1), cap(s1), ary)

	var s2 []int
	fmt.Printf("s2=%v, is nil=%t, len=%d, cap=%d\n", s2, s2 == nil, len(s2), cap(s2))

	// 切片自动扩容规则：容量小于1024，则每次*2；大于等于1024，每次增加1/4
	// 此时数组会还原，切片指向分配的新数组，与原数组无关联
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("s2=%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// 复制切片时，目标切片小于源切片长度，也可以复制且不报错，超过部分则截断
	s3 := []int{1, 2, 3, 4}
	s4 := make([]int, 3)
	copy(s4, s3)
	fmt.Println(s4)

	// 删除某个元素，使用append操作，尾段复制到头段，覆盖中段实现
	s5 := append(ary[:4], ary[5:]...)
	fmt.Println(s5)

	ary2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("ary2:%v\n", ary2)

	// 删除头元素
	s6 := ary2[1:]
	fmt.Println(s6)
	// 删除尾元素
	s7 := ary2[:len(ary2)-1]
	fmt.Println(s7)
}
