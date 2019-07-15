package main

import "fmt"

// 指针方法可以通过指针调用
// 值方法可以通过值调用
// 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
// 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
func main() {
	// 测试1：struct{}实现返回的是值
	l := List{1, 2, 3}

	// 方法要求参数是指针调用，所以加&
	AppendNumber(&l, 2)
	fmt.Println(l)

	// 方法要求参数是值调用，所以不加
	isLarge := IsLarge(l)
	fmt.Println(isLarge)

	fmt.Println("===================================")

	// 测试2：new返回的是指针
	var l2 = new(List)

	// 方法要求参数为指针调用，本身就是指针，所以调用
	AppendNumber(l2, 15)
	fmt.Println(l2)

	// 方法要求参数值调用，传递指针，因为可以自动解包
	isLarge2 := IsLarge(l2)
	fmt.Println(isLarge2)

	fmt.Println("===================================")

	course := Course{
		"Go语言开发",
		"杨明",
	}
	fmt.Println(course.String())
	fmt.Printf("%v\n", course.String())
}

func AppendNumber(appender *List, c int) {
	for i := 0; i < c; i++ {
		appender.Append(i)
	}
}

func IsLarge(l Len) bool {
	return l.Len() > 10
}

// 定义了一个List类型，实现了下面的两个接口

type List []int

func (l *List) Append(i int) {
	*l = append(*l, i)
}

func (l List) Len() int {
	return len(l)
}

// 定义了以下两个接口

type Appender interface {
	Append(int)
}

type Len interface {
	Len() int
}

// 系统接口测试
type Course struct {
	title   string
	teacher string
}

func (c *Course) String() string {
	return fmt.Sprintf("title:%s, teacher:%s", c.title, c.teacher)
}
