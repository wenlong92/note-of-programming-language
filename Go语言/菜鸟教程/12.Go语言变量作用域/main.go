package main

import "fmt"

var g int // 声明全局变量
var g1 int = 20

var a3 int = 20

func main() {
	// 作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围：Go语言中变量可以在三个地方声明
	// 1.函数内定义的变量称为局部变量
	// 2.函数外定义的变量称为全局变量
	// 3.函数定义中的变量称为形式参数

	// 1.局部变量
	var a1, b1, c1 int
	a1 = 10
	b1 = 20
	c1 = a1 + b1
	fmt.Printf("结果 a1 = %d, b1 = %d, c1 = %d \n", a1, b1, c1)

	// 2.全局变量：可以在整个包甚至外部包（被导出后）使用
	var a2, b2 int
	a2 = 10
	b2 = 20
	g = a2 + b2
	fmt.Printf("结果 a2 = %d, b2 = %d, g = %d\n", a2, b2, g)
	// 全局变量可以和局部变量重名，但函数内的局部变量会被优先考虑
	g1 = 10
	fmt.Printf("结果 g1 = %d\n", g1)

	// 3.形式参数：形式参数会作为函数的局部变量来使用
	var a3 = 10
	var b3 = 20
	var c3 = 0

	fmt.Printf("main()函数中 a3 = %d\n", a3)
	c3 = sum(a3, b3)
	fmt.Printf("main()函数 c3 = %d\n", c3)
}

func sum(a3, b3 int) int {
	fmt.Printf("sum()函数中 a3 = %d\n", a3)
	fmt.Printf("sum()函数中 b3 = %d\n", b3)
	return a3 + b3
}

// 不同类型的局部和全局变量默认值
// int  0, float32  0, pointer  nil
