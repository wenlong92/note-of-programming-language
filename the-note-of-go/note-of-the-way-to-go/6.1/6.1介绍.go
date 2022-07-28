package main
// import (
// 	"fmt"
// )

func main() {
	// 函数是基本的代码块
	// Go是编译型语言，所以函数编写顺序无关紧要

	// Go中有三种类型的函数
	// 1.普通的带有名字的函数
	// 2.匿名函数或者lambda函数
	// 3.方法(Methods)

	// 除了main()、init()，其他所有类型函数都可以有参数和返回值。
	// 函数参数、返回值以及它们的类型统称为函数签名

	// 下面是不正确写法：
	// func g()
	// {
	// }

	// 函数的正确写法：
	// func g() {
	// }

	// 函数被调用的基本格式：
	// pack1.Function(arg1, arg2, ..., argn)

	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")

	// 函数重载：编写多个同名函数，只要它们的形参或返回值不同
	// Go不允许函数重载

	// 函数是一等值，可以赋值给变量  add := binOp

	// 函数值之间可以相互比较，如果引用的是相同的函数或者都是nil，则认为它们是相同的函数
	// 函数不能在其他函数里声明(不能嵌套)，不过可以使用匿名函数

	// Go目前没有泛型概念，也就是不支持多种类型的函数。
	// 不过大部分情况可以通过接口(interface)，特别是空接口与类型选择与/或者通过使用反射(reflection)实现类似功能
	// 使用这些技术会导致代码更为复杂、性能低下
	// 在注重性能的场景下，最好为每个类型单独创建一个函数，代码可读性更强。
}

func greeting() {
	println("In greeting: Hi!!!!")
}