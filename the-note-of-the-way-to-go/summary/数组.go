package main

import (
	"fmt"
)

func main() {
	// 数组声明：要指定元素类型以及元素个数，元素类型要相同，例如：
	var balance [10] float32
	fmt.Println("balance:",balance)

	// 初始化：
	var balance0 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance0:",balance0)

	// 也可以通过字面量在声明的同时快速初始化：
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance1:",balance1)

	// 如果长度不确定，可以用...代替长度，编译器会根据个数自行判断数组长度，例如：
	var balance2 = [...]float32{2.3, 4.5, 5.6}
	fmt.Println("balance2:",balance2)
	balance3 := [...]float32{1.0, 2.1, 3.2, 4.3}
	fmt.Println("balance3:",balance3)

	// 如果设置了长度，可以通过索引初始化指定元素：
	balance4 := [3]int{0:3, 2:9}
	fmt.Println("balance4:",balance4)


	// 多维数组
	// 以下声明了三维的整型数组：
	var threedim [5][10][4]int
	fmt.Println("threedim:", threedim)

	// 二维数组
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println("Row1:", values[0])
	fmt.Println("Row2:", values[1])
	fmt.Println("第一个元素为", values[0][0])
	fmt.Println("values:", values)

	// 二维数组初始化，大多可通过大括号进行初始化，例如：
	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println("a:", a)


	// 创建空的二维数组
	animals := [][]string{}
	// 创建三一维数组，各数组长度不同
	row10 := []string{"fish", "shark", "eel"}
	row20 := []string{"bird"}
	row30 := []string{"lizard", "salamander"}
	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row10)
	animals = append(animals, row20)
	animals = append(animals, row30)
	fmt.Println(animals)


	// Go语言向函数传递数组
	// 向函数传递数组参数时，需要在函数定义时，声明形参为数组，有以下两种方式：
	// 第一种： void myFunction(param [10]int) { ... }
	// 第二种： void myFunction(param []int) { ... }

}