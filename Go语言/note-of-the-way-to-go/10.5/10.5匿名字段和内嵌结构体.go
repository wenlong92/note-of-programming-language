package main
import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	float64
	int     // anonymous field
	innerS  // anonymous field
}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	// 10.5.1定义
	// 结构体可以包含一个或多个匿名字段(内嵌字段)，即这些字段没有显式的名字，只有字段的类型时必须的，此时类型就是字段的名字
	// 匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体

	// 可以将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。
	// Go语言中的继承是通过内嵌或组合来实现的，可以说在Go语言中，相比较继承，组合更受青睐
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.float64 = 1.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.float64 is: %f\n", outer.float64)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 1.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:",outer2,"\n")
	// 通过类型 outer.int 的名字来获取存储在匿名字段中的数据，得出结论：在一个结构体中对于每一种数据类型只能有一个匿名字段

	// 10.5.2内嵌结构体
	// 结构体也是一种数据类型，也可以作为一个匿名字段使用，如上例中那样。
	// 外层结构通过outer.in1直接进入内层结构体的字段，内嵌结构体甚至可以来自其他包。
	// 内层结构体被简单的插入或者内嵌进外层结构体。
	// 这个简单的“继承”机制提供了一种方式，是的可以从另一个或一些类型继承部分或全部实现。
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)

	// 10.5.3命名冲突
	// 当两个字段拥有相同的名字(可能是继承来的名字)时：
	// 1.外层名字会覆盖内层名字(但两者的内存空间都保留)，这提供了一种重载字段或方法的方式；
	// 2.如果相同的名字在同一级别出现两次，如果这个名字被程序使用了，将会引发一个错误(不使用没关系)。没有办法解决这种问题引起的二义性，必须由程序员自己修正。
	type A1 struct {a int}
	type B1 struct {a, b int}
	type C1 struct {A; B}
	var c1 C1
	// 此时使用c1.a是错误的，会导致编译器错误
	type D1 struct {B1; b float32}
	var d1 D1
	// 此时使用d1.b没问题，他是float32，而不是B1 的b，如果想得到内层的b，可通过d1.B1.b 
}