package main

import "fmt"

func main() {
	// 以下有待仔细研究，new返回什么类型，Type{}返回什么类型？为什么Type{}返回的实例需要使用引用传递

	myPhone := new(Honor)
	myPhone.Name = "Honor"
	ShowFunctions(myPhone)

	fmt.Println("=============")

	myPhone2 := Honor{
		"Honor 10",
	}
	ShowFunctions(&myPhone2)

	fmt.Println("***************************")

	GetAnyType(123)
	GetAnyType("123")
	GetAnyType(myPhone)
	GetAnyType(*myPhone)

	fmt.Println("***************************")

	q := Queue{1, 2}
	q.Push(3)
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)

	q.Push("Chris")
	fmt.Println(q)
}

// 空接口，用interface{}表示任何一个数据类型
type Queue []interface{}

func (q *Queue) Push(element interface{}) {
	*q = append(*q, element)
}

func (q *Queue) Pop() interface{} {
	// 这里为什么要用()将*q括起来
	// 因为*q为Queue类型，取索引是数组的行为，所以需要用括号括起来做强制转换
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

/////////////////

// 使用type-switch获取具体数据类型 
func GetAnyType(any interface{}) {
	switch t := any.(type) {
	case int:
		fmt.Println("any is int type")
	case string:
		fmt.Println("any is string type")
	default:
		fmt.Printf("any is %T\n", t)
	}
}

/////////////////

type Phone interface {
	Call(string)
	SMS(string)
}

type Camera interface {
	Take() string
}

type SmartPhone interface {
	Phone
	Camera
	Download(string) string
}

type Honor struct {
	Name string
}

func (phone *Honor) Call(number string) {
	fmt.Println(phone.Name, " call to ", number)
}

func (phone *Honor) SMS(msg string) {
	fmt.Println(phone.Name, " show msg:", msg)
}

func (phone *Honor) Take() string {
	return "avatar.jpg"
}

func (phone *Honor) Download(url string) string {
	return "wechat.pkg"
}

func ShowFunctions(sp SmartPhone) {
	sp.Call("10000")
	sp.SMS("i receive a msg")

	photo := sp.Take()
	fmt.Printf("photo:%s\n", photo)

	pkg := sp.Download("http://ichris.info/app.pkg")
	fmt.Printf("pkg:%s", pkg)
}
