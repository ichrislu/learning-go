package main

import "fmt"

// Book struct
type Book struct {
	id     int
	title  string
	author string
}

// GetAuthor 结构体方法
func (book *Book) String() string {
	return fmt.Sprintf("id=%d, title=%s, author=%s", book.id, book.title, book.author)
}

// GetAuthor 结构体方法
func (book *Book) GetAuthor() string {
	return book.author
}

// 当结构体小写时，可以强制使用包内定义好的工厂方法，统一结体体实例化方法

// NewBook 工厂方法初始化Book
func NewBook(id int, title, author string) *Book {
	return &Book{id, title, author}
}

func main() {
	var book1 *Book
	book1 = new(Book)
	book1.id = 1
	book1.title = "golang in action"
	book1.author = " some body"
	fmt.Println(book1)

	// 直接初始化
	book2 := Book{
		id:     2,
		title:  "kubernetes in action",
		author: "some one",
	}
	fmt.Println(book2)

	// 省略字段直接初始化，但必须按照定义顺序
	book3 := Book{
		3,
		"spring boot in action",
		"spring ico.",
	}
	fmt.Println(book3)

	// 普通参数传递，是值传递，不影响外面的book3实例
	change(book3)
	fmt.Println(book3.String())

	// 指针参数传递，影响外面的book3实例
	changePtr(&book3)
	fmt.Println(book3.String())

	fmt.Println(book3.GetAuthor())

	book4 := NewBook(4, "测试书", "超哥")
	fmt.Println(book4)
}

func change(book Book) {
	book.title += " changed"
}

func changePtr(book *Book) {
	book.title += " changed by PTR"
}
