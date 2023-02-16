package main

// 声明变量的一般形式为： var identifier type
// 变量名由字母、数字、下划线组成，首个字符不能是数字

import "fmt"

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 如果没有初始化，变量默认为零值
	var d int
	fmt.Println(d)

	// 根据值自行判断变量类型
	var e = true
	fmt.Println(e)

	// 简写形式
	f := false
	fmt.Println(f)

	// int float bool string这些基本类型属于值类型

	// 可通过 &i 获取变量内存地址，称之为指针
	// 更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存

	// 多变量可以在同一行进行赋值
	var aaa, bbb int
	var ccc string
	aaa, bbb, ccc = 1, 2, "hello"
	print(aaa, bbb, ccc)

	// 如果想交换两个变量的值，可以用 a, b = b, a，两个变量类型需相同

	// 空白标识符 _ 用于抛弃值，它实际上是一个只写变量

	// 并行赋值也被用于当一个函数返回多个返回值时，例如 val, err = Func1(var1)

}

// 全局变量允许声明但不使用，局部变量不行
// 同一类型的多个变量可以声明在同一行
var aa, bb, cc int
