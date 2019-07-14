package main

import (
	"fmt"

	"../ch17/book"
)

func main() {
	book1 := book.NewBook(1, "设计模式", "大神")
	fmt.Println(book1)

	book1.SetTitle("设计模式2")
	fmt.Println(book1)

	book1.SetAuthor("超级大神")
	fmt.Println(book1)

	// 反射，这里还没有搞太明白
	book.RefTag(*book1, 1)

	techBk := book.InitTechBook(2, "Python in action", "Python大神", "二楼中间", 100)
	fmt.Println(techBk)
	fmt.Printf("isdn:%d, title:%s, author:%s, pos:%s, count:%d\n",
		techBk.GetId(), techBk.GetTitle(), techBk.GetAuthor(), techBk.GetPos(), techBk.GetCount())

	// 内嵌结构体
	a := book.A{A1: "a.a1"}
	b := book.B{A1: "b.a1", B1: "b.b1"}
	c := book.C{A: a, B: b}

	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("c:%v\n", c)
	fmt.Printf("同名属性:%v\n", c.A.A1) // 同名属性，需要指明类型
	fmt.Printf("同名属性:%v\n", c.B.A1)
}
