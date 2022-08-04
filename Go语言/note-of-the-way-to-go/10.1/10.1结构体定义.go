package main

import (
	"strings"
	"fmt"
)

func main() {
	// 定义方式如下：

	// type identifier struct {
	// 	field1 type1
	// 	field2 type2
	// 	...
	// }

	// 或者 type T struct {a, b int}，更适用于简单的结构体

	// 结构体里的字段都有名字，例如field1、field2等，若字段在代码中不会被用到，可以命名为_

	// 结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或接口
	// 可以声明结构体类型的一个变量，然后给它的字段赋值,例如：

	// var s T
	// s.a = 5
	// s.b = 8

	// 数组可看做是一种结构体类型，不过它使用下标而不是具名的字段

	// 使用new函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)
	// 这条语句的惯用方法是：t := new(T),变量t是一个指向T的指针，此时结构体字段的值是它们所属类型的零值

	// 声明 var t T 也会给t分配内存并零值化内存，但这时t是类型T
	// 这两种方式中，t通常被称作类型T的一个实例或对象
	
	type struct1 struct {
		i1 int
		f1 float32
		str string
	}
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	// type myStruct struct { i int }
	// var v myStruct       // v 是结构体类型变量
	// var p *myStruct      // p 是指向一个结构体类型变量的指针
	// v.i
	// p.i

	// 初始化一个结构体实例更简短和惯用的方式如下：
	// ms := &struct1{10, 15.5, "Chris"}   // 此时 ms 的类型时 *struct1
	// 或者：
	// var ms struct1
	// ms = struct1{10, 15.5, "Chris"}
	// 混合字面量语法 &struct1{a, b, c} 是一种简写，底层仍调用 new()，这里值的顺序必须按照字段顺序来写

	// 时间间隔是使用结构体的一个典型例子：
	type Interval struct {
		start int
		end int
	}
	// 初始化方式：
	// intr1 := Interval{0, 3}            //(A)
	// intr2 := Interval{end:5, start:1}  //(B)
	// intr3 := Interval{end:5}           //(C)

	// 结构体类型和字段的命名遵循可见性规则(4.2)，一个导出的结构体类型中有些字段是导出的，另一些不是，这是可能的

	// 下面Person结构体的例子，展示了三种不同的方式：
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2-struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	// (*pers2).lastName = "demo"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3-struct as a literal
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	// 结构体的内存布局
	// 结构体和它所包含的数据在内存中是以连续快的形式存在的，即使结构体中嵌套有其他的结构体，在性能上带来了很大的优势


	// 递归结构体
	// 结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素时特别有用，此时节点包含指向临近节点的链接


	// 结构体转换
	// Go中的类型转换遵循严格的规则。当为结构体定义了一个alias类型时，此结构体类型和它的alias类型都有相同的底层类型
	// 可以如下例中相互转换，例如：
	type number struct {
		f float32
	}
	type nr number
	a := number{5.0}
	b := nr{5.0}
	var c = number(b)
	fmt.Println(a,b,c)
}
type Person struct {
	firstName string 
	lastName string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}