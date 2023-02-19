package main

// 常量是一个简单值的标识符，运行时不会被修改的量
// 常量的数据类型只能是布尔型、数字型（整数型、浮点型、复数）和字符串型

// 显示类型定义：const b string = "abc"
// 隐式类型定义：const b = "abc"

// 多个相同类型：const a, b = "abc", "ade"

import (
	"fmt"
	"unsafe"
)

func main() {
	const LENGTH int = 10
	const WIDTH = 5
	var area int
	const a2, b2, c2 = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a2, b2, c2)

	// 常量还可以用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// 常量可以用len() cap() unsafe.Sizeof()函数表达式的值，必须是内置函数
	const (
		a1 = "abc"
		b1 = len(a1)
		c1 = unsafe.Sizeof(a1)
	)
	println(a1, b1, c1)

	// iota用法:特殊常量，可以认为是一个可以被编译器修改的常量
	// iota在constructor关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次
	// （iota可理解为const语句块中的行索引）

	// iota可以被用作枚举值
	//const (
	//   a2 = iota
	//   b2 = iota
	//   c2 = iota
	//)
	//第一个iota等于0，每当iota在新的一行被使用时，它的值都会自动加1；所以a=0,b=1,c=2可简写为：
	//const (
	//   a = iota
	//   b
	//   c
	//)

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // ha 遇到这种独立值，显示独立值不显示iota，但此时iota继续+1，所以iota为3
		e        // ha 如果从上面独立值到这，没有iota干扰，那么这里等于上面的独立值，所以显示为ha，但iota继续+1，iota为4
		f = 100  // 100 遇到独立值，显示为100，iota+1，iota为5
		g        // 100 与上面独立值相同，显示为100，iota+1，iota为6
		h = iota // 7   这里开始将iota显示出来，所以iota+1后为7
		i        // 8   iota+1位8
		// 这里的重点是将iota看做行索引
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		j = 1 << iota // 此时iota为0，左移0位，所以j不变，仍为1
		k = 3 << iota // 此时iota为1，左移1位，3的二进制表示为011,所以左移1位后二进制为110，变为6
		l             // 可看做l=3<<iota，此时iota为2，左移2位，为1100，此时为12
		m             // 可看做m=3<<iota，此时iota为3，左移3位，为11000，此时为24
	)
	fmt.Println(j, k, l, m)
}
