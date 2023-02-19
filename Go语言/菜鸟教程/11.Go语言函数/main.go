package main

import "fmt"

func main() {
	// 函数是基本的代码块，用于执行一个任务
	// GO语言中最少有一个main() 函数
	// 可通过函数来划分不同功能，逻辑上每个函数执行的是指定任务
	// 函数声明告诉了编译器函数的名称，返回类型和参数
	// Go语言标准库提供了多种可动用的内置函数。如：len()函数可接受不同类型参数并返回该类型的长度。如果传入的是字符串则返回字符串的长度，如果传入的是数组则返回数组中包含元素的个数

	// 函数调用
	var a, b int = 100, 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值是： %d \n", ret)

	// 返回多个值
	a1, b1 := swap("Google", "Runoob")
	fmt.Println(a1, b1)
}

// 函数实例
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
