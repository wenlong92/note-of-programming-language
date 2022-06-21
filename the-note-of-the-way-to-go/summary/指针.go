package main

import (
	"fmt"
)

func main() {
	// 变量是一种使用方便的占位符，用于引用计算机内存地址
	// Go语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	var a int = 10
	fmt.Printf("地址：%x \n", &a)


	// 什么是指针：一个指针变量指向了一个值的内存地址
	// 类似于变量和常量，在使用指针前需要声明指针，声明格式如下：
	// var ip *int
	// var fp *float32

	// 指针使用流程：定义指针变量、为指针变量赋值、访问指针变量中指向地址的值
	a1 := 10       //声明实际变量
	var ip1 *int   //声明指针变量
	ip1 = &a1      //指针变量的存储地址
	fmt.Printf("a1变量的地址是：%x\n", &a1)
	fmt.Printf("ip1变量存储的地址：%x\n", ip1)
	fmt.Printf("*ip1变量的值是：%d\n", *ip1)

	// 空指针：定义后没有分配任何变量
	var ptr *int
	fmt.Printf("ptr 的值为：%x\n", ptr)

}