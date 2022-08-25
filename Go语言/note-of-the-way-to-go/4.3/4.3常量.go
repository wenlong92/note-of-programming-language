package main

import (
	"fmt"
)

func main() {
	// 存储在常量中的数据类型只能是布尔型、数字型(整数型、浮点型和复数)、字符串型
	
	const pi = 3.1415926
	
	const b string = "abc"
	
	// 常量的值必须是在编译时就能够确定的；可以在赋值表达式中设计计算过程，但是所有用于计算的值必须在编译期间就能获得
	// 正确的做法：const c1 = 2/3
	// 错误的做法：const c2 = getNumber()   
	// 因为在编译期间自定义函数属于未知，所以无法用于常量的赋值，但内置函数可以使用，如：len()
	
	// 并行赋值
	const beef, two, c = "eat", 2, "veg"
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	
	// 枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	
	// iota可以被用作枚举值
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	fmt.Println("a1:", a1)
	fmt.Println("b1:", b1)
	fmt.Println("c1:", c1)

	// 第一个iota等于0，每当iota在新的一行被使用时，它的值都会自动加1，并且没有赋值的常量默认会应用上一行的赋值表达式
	const (
		a2 = iota
		b2
		c2
		d2 = 5
		e2
	)
	fmt.Println("a2:", a2)
	fmt.Println("b2:", b2)
	fmt.Println("c2:", c2)
	fmt.Println("d2:", d2)
	fmt.Println("e2:", e2)


	// 无法在程序运行中修改常量的值，否则会引发编译错误。

}

