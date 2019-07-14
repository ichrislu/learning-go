package book

type techBook struct {
	pos   string
	count int
	// 覆盖book中title
	title string
	book
}

func InitTechBook(id int, title, author string, _pos string, _count int) *techBook {
	bk := NewBook(id, title, author)
	return &techBook{
		pos:   _pos,
		count: _count,
		title: "覆盖基类",
		book:  *bk,
	}
}

// 这里为什么不需要tb.book.id
func (tb *techBook) GetId() int {
	return tb.id
}

func (tb *techBook) GetTitle() string {
	return tb.title
}

func (tb *techBook) GetAuthor() string {
	return tb.author
}

func (tb *techBook) GetPos() string {
	return tb.pos
}

func (tb *techBook) GetCount() int {
	return tb.count
}
