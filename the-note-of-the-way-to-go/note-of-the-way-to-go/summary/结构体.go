package main

import (
	"fmt"
)

func main() {
	// Go语言中数组可以存储同一类型的数据，在结构体中可以为不同项定义不同的数据类型
	// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
	// 结构如下：

	// type struct_variable_type struct {
	// 	member definition
	// 	member definition
	// 	...
	// 	member definition
	// }

	// 定义结构体类型后，可用于变量的声明，语法格式如下：
	// variable_name := structure_variable_type { value1, value2...value}  
	//或者
	// variable_name := structure_variable_tyep { key1: value1, key2: value2..., keyn: valuen}

	// 例子：
	fmt.Println(Books{"bookTitle", "bookAuthor", "bookSubject", 123456})
	// 也可以用key => value 格式
	fmt.Println(Books{title:"Go title", author:"Go author", subject:"Go subject", book_id:45679})
	// 忽略的字段为0或空
	fmt.Println(Books{title:"Go语言", author: "Bill"})

	// 访问结构体成员用 . 操作符：结构体.成员名
	var Book1 Books
	var Book2 Books
	Book1.title = "book1Title"
	Book1.author = "book1Author"
	Book1.subject = "book1Subject"
	Book1.book_id = 111111

	Book2.title = "book2Title"
	Book2.author = "book2Author"
	Book2.subject = "book2Subject"
	Book2.book_id = 2222222

	fmt.Println(Book1)
	fmt.Println(Book2)


	// 结构体作为函数参数
	printBook(Book1)


	// 结构体指针
	// 可以定义指向结构体的指针，类似于其他指针变量，格式如下：
	var struct_pointer *Books
	// 查看结构体变量地址，还是用 &
	struct_pointer = &Book1
	fmt.Println(struct_pointer)

	printBook1(&Book1)

}
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBook(book Books) {
	fmt.Println("printBook:", book)
}

func printBook1(book *Books) {
	fmt.Println("printBook1:", book)
}