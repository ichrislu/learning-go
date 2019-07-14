package book

import (
	"fmt"
	"reflect"
)

// Book struct
type book struct {
	id     int    "ISBN编号"
	title  string "标题"
	author string "作者"
}

// NewBook 构造函数
func NewBook(id int, title, author string) *book {
	return &book{id, title, author}
}

func (book *book) SetID(id int) {
	book.id = id
}

func (book *book) SetTitle(title string) {
	book.title = title
}

func (book *book) SetAuthor(author string) {
	book.author = author
}

// RefTag 反射构建Tag
func RefTag(book book, idx int) {
	bType := reflect.TypeOf(book)
	field := bType.Field(idx)
	fmt.Printf("%v\n", field.Tag)
}
