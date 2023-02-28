package main

import "fmt"

// 定义结构体
type Circle struct {
	radius float64
}

func main() {
	// Go语言中同时有函数和方法。
	// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个追或者是一个指针
	// 所有给定类型的方法属于该类型的方法集

	// func (variable_name variable_data_type) function_name() [return_type] {
	//   函数体
	//}

	// 结构体类型和该类型的一个方法
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

// 该method属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
