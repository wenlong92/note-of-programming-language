package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// Go语言中结构体中可以为不同项定义不同的数据类型
	// 结构体是由一系列具有想同类型或不同类型的数据构成的数据集合
	fmt.Println(Books{"Go语言", "www.runoob.com", "Go语言教程", 12121212})
	fmt.Println(Books{title: "Go语言", author: "www.runoob.com", subject: "Go语言教程", book_id: 12121212})
	fmt.Println(Books{title: "Go语言", author: "www.runoob.com"})

	// 访问结构体成员
	var Book1 Books
	var Book2 Books
	Book1.title = "title1"
	Book1.author = "author1"
	Book1.subject = "subject1"
	Book1.book_id = 1

	Book2.title = "title2"
	Book2.author = "author2"
	Book2.subject = "subject2"
	Book2.book_id = 2

	// 结构体作为函数参数
	printBook(Book1)
	printBook(Book2)

	// 结构体指针
	printBook1(&Book1)
	printBook1(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book title : %d\n", book.book_id)
}
func printBook1(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book title : %d\n", book.book_id)
}
