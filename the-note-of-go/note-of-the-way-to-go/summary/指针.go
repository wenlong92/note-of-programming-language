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


	// Go语言指针数组
	// 先定义一个长度为3的整型数组
	const MAX int = 3
	a2 := []int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a2[%d] = %d\n", i, a2[i])
	}
	// 如果需要保存数组，需要用到指针
	var ptr1 [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr1[i] = &a2[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a2[%d] = %d\n", i, *ptr1[i])
	}


	// Go语言指向指针的指针
	// 如果一个指针变量存放的又是另一个指针变量的地址，称这个指针变量为指向指针的指针，声明格式如下：
	// var ptr2 **int

	var a3 int
	var pptr **int
	var ptr3 *int

	a3 = 1000
	ptr3 = & a3
	pptr = &ptr3
	fmt.Printf("变量 a3 = %d\n", a3)
	fmt.Printf("指针变量 *ptr3 = %d\n", *ptr3)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)


	// Go语言指针作为函数参数
	// 只需要在函数定义的参数上设置为指针类型即可
	demo1 := 100
	demo2 := 200
	swap(&demo1, &demo2)
	fmt.Printf("交换后的demo1:%d\n",demo1)
	fmt.Printf("交换后的demo2:%d\n",demo2)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}